package main

import (
	"fmt"
	"io"
	"log"
	"os"

	uncp "github.com/ray-g/gopl/ch10/ex10.02/uncompressor"
	_ "github.com/ray-g/gopl/ch10/ex10.02/uncompressor/tar"
	_ "github.com/ray-g/gopl/ch10/ex10.02/uncompressor/zip"
)

func printArchive(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := uncp.Open(f)
	if err != nil {
		return fmt.Errorf("open archive reader: %s", err)
	}
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		return fmt.Errorf("printing: %s", err)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: extract FILE ...")
	}
	exitCode := 0
	for _, filename := range os.Args[1:] {
		err := printArchive(filename)
		if err != nil {
			log.Print(err)
			exitCode = 2
		}
	}
	os.Exit(exitCode)
}
