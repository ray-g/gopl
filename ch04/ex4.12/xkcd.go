package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type WordIndex map[string]map[int]bool
type NumIndex map[int]Comic

type Comic struct {
	Num              int
	Year, Month, Day string
	Title            string
	Transcript       string
	Alt              string
	Img              string // url
}

const usage = `xkcd get N
xkcd index OUTPUT_FILE
xkcd search INDEX_FILE QUERY`

func usageDie() {
	fmt.Println(usage)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	cmd := os.Args[1]
	switch cmd {
	case "get":
		if len(os.Args) != 3 {
			usageDie()
		}
		n, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "N (%s) must be an int", os.Args[1])
			usageDie()
		}
		comic, err := getComic(n)
		if err != nil {
			log.Fatal("Error getting comic", err)
		}
		fmt.Println(comic)
	case "index":
		if len(os.Args) != 3 {
			usageDie()
		}
		err := index(os.Args[2])
		if err != nil {
			log.Fatal("Error serializing indexes", err)
		}
	case "search":
		if len(os.Args) != 4 {
			usageDie()
		}
		filename := os.Args[2]
		query := os.Args[3]
		err := search(query, filename)
		if err != nil {
			log.Fatal("Error searching index", err)
		}
	default:
		usageDie()
	}
}

func getComic(n int) (Comic, error) {
	var comic Comic
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", n)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return comic, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("can't get comic %d: %s", n, resp.Status)
	}
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}
	return comic, nil
}

func index(filename string) error {
	nworkers := 20
	done := make(chan int, 0)
	comicChan, err := getComics(nworkers, done)
	if err != nil {
		return err
	}

	go func() {
		for i := 0; i < nworkers; i++ {
			<-done
		}
		close(done)
		close(comicChan)
	}()

	indexComicsDealer(comicChan, filename)

	return nil
}

func getComics(nworkers int, done chan int) (chan Comic, error) {
	max, err := getComicCount()
	if err != nil {
		return nil, err
	}
	fmt.Println("max", max)

	comics := make(chan Comic, 5*nworkers)
	comicNums := make(chan int, 1*nworkers)

	for i := 0; i < nworkers; i++ {
		go fetcher(comicNums, comics, done)
	}

	go dispatcher(comicNums, max)

	return comics, nil
}

func indexComicsDealer(comicChan chan Comic, filename string) {
	wordIndex, numIndex := indexComics(comicChan)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	// fmt.Println(wordIndex)
	err = enc.Encode(wordIndex)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = enc.Encode(numIndex)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func indexComics(comics chan Comic) (WordIndex, NumIndex) {
	wordIndex := make(WordIndex)
	numIndex := make(NumIndex)
	for comic := range comics {
		numIndex[comic.Num] = comic
		scanner := bufio.NewScanner(strings.NewReader(comic.Transcript))
		fmt.Printf("Dealing with comic %d\n", comic.Num)
		scanner.Split(ScanWords)
		for scanner.Scan() {
			token := strings.ToLower(scanner.Text())
			if _, ok := wordIndex[token]; !ok {
				wordIndex[token] = make(map[int]bool, 1)
			}
			wordIndex[token][comic.Num] = true
		}
		// fmt.Printf("Done with comic %d\n", comic.Num)
	}
	// fmt.Println("Done with all comic")
	return wordIndex, numIndex
}

func fetcher(comicNums chan int, comics chan Comic, done chan int) {
	for n := range comicNums {
		// fmt.Println("Getting", n)
		comic, err := getComic(n)
		if err != nil {
			log.Printf("Can't get comic %d: %s", n, err)
			continue
		}
		comics <- comic
	}
	fmt.Println("done")
	done <- 1
}

func dispatcher(comicNums chan int, max int) {
	for i := 1; i <= max; i++ {
		comicNums <- i
		// fmt.Println("added", i)
	}
	close(comicNums)
	fmt.Println("closed nums")
}

func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	i := 0
	start := 0
	stop := 0
	for i < len(data) {
		r, size := utf8.DecodeRune(data[i:])
		i += size
		if unicode.IsLetter(r) {
			start = i - size
			break
		}
	}
	for i < len(data) {
		r, size := utf8.DecodeRune(data[i:])
		i += size
		if !unicode.IsLetter(r) {
			stop = i - size
			break
		}
	}
	if stop > start {
		token = data[start:stop]
	}
	return i, token, nil
}

func readIndex(filename string) (WordIndex, NumIndex, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	dec := gob.NewDecoder(file)
	var wordIndex WordIndex
	var numIndex NumIndex
	err = dec.Decode(&wordIndex)
	if err != nil {
		return nil, nil, err
	}
	dec.Decode(&numIndex)
	if err != nil {
		return nil, nil, err
	}
	return wordIndex, numIndex, nil
}

func comicsContainingWords(words []string, wordIndex WordIndex, numIndex NumIndex) []Comic {
	found := make(map[int]int) // comic Num -> count words found
	comics := make([]Comic, 0)
	for _, word := range words {
		for num := range wordIndex[word] {
			found[num]++
		}
	}
	for num, nfound := range found {
		if nfound == len(words) {
			comics = append(comics, numIndex[num])
		}
	}
	return comics
}

func search(query string, filename string) error {
	wordIndex, numIndex, err := readIndex(filename)
	if err != nil {
		return err
	}
	comics := comicsContainingWords(strings.Fields(query), wordIndex, numIndex)
	for _, comic := range comics {
		fmt.Printf("%+v\n\n", comic)
	}
	return nil
}

func getComicCount() (int, error) {
	resp, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("can't get main page: %s", resp.Status)
	}
	var comic Comic
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return 0, err
	}
	return comic.Num, nil
}
