package memo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

//!+httpRequestBody
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//!-httpRequestBody

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, done <-chan bool) (interface{}, error)
}

/*
//!+seq
	m := memo.New(httpGetBody)
//!-seq
*/

func Sequential(t *testing.T, m M) {
	//!+seq
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, nil)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
	//!-seq
}

/*
//!+conc
	m := memo.New(httpGetBody)
//!-conc
*/

func Concurrent(t *testing.T, m M) {
	//!+conc
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, nil)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
	//!-conc
}

func Test(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	Concurrent(t, m)
}

func TestCancel(t *testing.T) {
	m := New(httpGetBody)
	defer m.Close()
	key := "https://golang.org"
	wg1 := &sync.WaitGroup{}
	wg1.Add(1)
	go func() {
		v, err := m.Get(key, nil)
		wg1.Done()
		if v == nil {
			t.Errorf("got %v, %v; want %v, %v", v, err, nil)
		}
	}()
	wg1.Wait()

	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		done := make(chan bool)
		close(done)
		v, err := m.Get(key, done)
		if v != nil || err == nil {
			t.Errorf("got %v, %v; want %v, %v", v, err, nil, "cancled")
		}
		wg2.Done()
	}()
	wg2.Wait()
}
