// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	wg := sync.WaitGroup{}
	defer func() {
		wg.Wait()
		// NOTE: ignoring potential errors from input.Err()
		c.Close()
	}()
	timeout := 10 * time.Second
	timer := time.NewTimer(timeout)
	inputs := make(chan string)
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			inputs <- input.Text()
		}
		if input.Err() != nil {
			log.Println("scan:", input.Err())
		}
	}()

	for {
		select {
		case input := <-inputs:
			timer.Reset(timeout)
			wg.Add(1)
			go func() {
				defer wg.Done()
				echo(c, input, 1*time.Second)
			}()
		case <-timer.C:
			return
		}

	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
