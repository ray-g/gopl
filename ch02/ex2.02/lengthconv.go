package main

import (
	"fmt"
)

// Feet FT
type Feet float64

// Meter M
type Meter float64

// FToMFactor factor of Feet to Meter, 1 Meter = 3.28 Feet
const FToMFactor = 3.28

func (f Feet) String() string  { return fmt.Sprintf("%.2f FT", f) }
func (m Meter) String() string { return fmt.Sprintf("%.2f M", m) }

// FToM converts Feet to Meter
func FToM(f Feet) Meter {
	return Meter(f / FToMFactor)
}

// MToF converts Meter to Feet
func MToF(m Meter) Feet {
	return Feet(m * FToMFactor)
}
