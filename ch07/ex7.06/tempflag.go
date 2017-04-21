// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var stdout io.Writer = os.Stdout
var stderr io.Writer = os.Stderr

//!+
var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Fprintln(stdout, *temp)
}

//!-

// Package tempconv performs Celsius and Fahrenheit temperature computations.
// package tempconv

// import (
// 	"flag"
// 	"fmt"
// )

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
	DeltaCK       Celsius = -273.15
	DeltaKC       Kelvin  = 273.15
)

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c - DeltaCK) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - DeltaKC) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

func (c Celsius) String() string    { return fmt.Sprintf("%.2f°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.2f°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%.2f°K", k) }

/*
//!+flagvalue
package flag
// Value is the interface to the value stored in a flag.
type Value interface {
	String() string
	Set(string) error
}
//!-flagvalue
*/

//!+celsiusFlag
// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

//!-celsiusFlag

//!+CelsiusFlag

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

//!-CelsiusFlag
