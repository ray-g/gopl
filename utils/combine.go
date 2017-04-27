// Generage Exercises.md
// run 'go run utils/combine.go -o Exercises.md' under root
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var oFlag = flag.String("o", "Exercises.md", "out file of combined exercises.")

type StringSlice []string

func (ss StringSlice) Len() int      { return len(ss) }
func (ss StringSlice) Swap(i, j int) { ss[i], ss[j] = ss[j], ss[i] }
func (ss StringSlice) Less(i, j int) bool {
	// all should be like this: ch01/ex1.03
	ei := strings.Split(ss[i], "ex")
	vi, _ := strconv.ParseFloat(ei[1], 64)

	ej := strings.Split(ss[j], "ex")
	vj, _ := strconv.ParseFloat(ej[1], 64)

	return vi < vj
}

var ReadMeDirs []string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()

	f, err := os.Create(*oFlag)
	check(err)
	defer f.Close()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg)
	}

	wg.Wait()

	sort.Sort(StringSlice(ReadMeDirs))

	fmt.Fprintln(f, "# Exercises\n")
	for _, dir := range ReadMeDirs {
		func() {
			r, err := os.Open(filepath.Join(dir, "README.md"))
			check(err)
			defer r.Close()

			rd := bufio.NewReader(r)
			first := true
			for {
				line, isPrefix, err := rd.ReadLine()
				if err != nil {
					if err == io.EOF {
						break
					}
					check(err)
					return
				}
				if first {
					// # Exercise ...
					title := fmt.Sprintf("\n## [%s](%s)", line[2:], dir)
					f.Write([]byte(title))
					first = false
				} else {
					f.Write(line)
				}
				if !isPrefix {
					f.Write([]byte("\r\n"))
				}
			}
		}()
	}
}

func walkDir(dir string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, wg)
		} else {
			if entry.Name() == "README.md" && dir != "." {
				ReadMeDirs = append(ReadMeDirs, dir)
			}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
