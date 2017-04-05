// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var stdout io.Writer = os.Stdout
var done = make(chan struct{})

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	waitForCancel()

	for range os.Args[1:] {
		fmt.Fprintln(stdout, <-ch) // receive from channel ch
	}
	fmt.Fprintf(stdout, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func waitForCancel() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			close(done)
			return
		}
	}()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	req.Cancel = done
	fmt.Println("requesting", url)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
