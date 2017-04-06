package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	tempconv "github.com/ray-g/gopl/ch02/ex2.01"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr
var stdin io.Reader = os.Stdin

func tempUnit(n float64) string {
	f := tempconv.Fahrenheit(n)
	c := tempconv.Celsius(n)
	return fmt.Sprintf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f),
		c, tempconv.CToF(c),
	)
}

func lenUnit(n float64) string {
	f := Feet(n)
	m := Meter(n)
	return fmt.Sprintf("%s = %s, %s = %s\n",
		f, FToM(f),
		m, MToF(m),
	)
}

func weightUnit(n float64) string {
	p := Pound(n)
	kg := Kilogram(n)
	return fmt.Sprintf("%s = %s, %s = %s\n",
		p, PToKG(p),
		kg, KGToP(kg),
	)
}

func allUnits(n float64) string {
	return tempUnit(n) + lenUnit(n) + weightUnit(n)
}

func printAll(nums []string) {
	for _, num := range nums[0:] {
		n, err := strconv.ParseFloat(num, 64)
		if err != nil {
			fmt.Fprintf(stderr, "unitconv: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(stdout, allUnits(n))
		fmt.Fprintf(stdout, "----------\n")
	}
}

func main() {
	if len(os.Args) == 1 {
		scanner := bufio.NewScanner(stdin)
		for scanner.Scan() {
			printAll(strings.Split(scanner.Text(), " "))
		}
		return
	}
	printAll(os.Args[1:])
}
