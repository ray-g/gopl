// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var stdout io.Writer = os.Stdout
var fileCount = 0

const fileBase = "fetch_test_result_"
const fileExt = ".html"

func main() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		go fetch(i, url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Fprintln(stdout, <-ch) // receive from channel ch
	}
	fmt.Fprintf(stdout, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(index int, url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	fileName := fileBase + strconv.Itoa(index) + fileExt

	file, err := os.Create(fileName)
	if err != nil {
		ch <- fmt.Sprintf("Failed to create file. error: %v", err)
	}
	defer file.Close()

	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
