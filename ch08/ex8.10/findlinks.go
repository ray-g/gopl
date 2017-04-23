package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/net/html"
)

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	cancel := make(chan struct{})    // for cancel all http.Request.

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()
	wg := sync.WaitGroup{}

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(num int) {
			func() { // To like a Label.
				for {
					select {
					case link := <-unseenLinks:
						foundLinks := crawl(link, cancel)
						log.Printf("Got link in %s\n", link)
						go func() {
							select {
							case <-cancel: // Cancel signal.
							default:
								worklist <- foundLinks
							}
						}()
					case <-cancel: // Cancel signal.
						return // Jump like a label.
					}
				}
			}()
			fmt.Printf("Canceled in crawler %d\n", num)
			wg.Done()
		}(i)
	}

	go func() { // Cancel timer
		time.Sleep(1 * time.Second)
		log.Printf("\nCanceled\n")
		close(cancel) // Cancel all http requests.
	}()

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	func() {
		for list := range worklist {
			for _, link := range list {
				select {
				case <-cancel:
					//close(unseenLinks) // End unseenLinks range.
					return
				default:
					if !seen[link] {
						seen[link] = true
						unseenLinks <- link
					}
				}
			}
		}
	}()
	wg.Wait() // Exit after other goroutines.
}

func crawl(url string, cancel <-chan struct{}) []string {
	list, err := Extract(url, cancel)
	if err != nil {
		log.Print(err)
	}
	return list
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string, cancel <-chan struct{}) ([]string, error) {
	// Create request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []string{}, err
	}

	// Set channel for cancel request.
	req.Cancel = cancel

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
