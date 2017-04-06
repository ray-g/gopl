package main

import (
	"fmt"
)

// Pound P
type Pound float64

// Kilogram KG
type Kilogram float64

// PToKGFactor factor of Pounds to Kilograms 1 KG = 2.2046 Pounds
const PToKGFactor = 2.2046

func (p Pound) String() string     { return fmt.Sprintf("%.2f P", p) }
func (kg Kilogram) String() string { return fmt.Sprintf("%.2f KG", kg) }

// PToKG converts Pounds to Kilograms
func PToKG(p Pound) Kilogram {
	return Kilogram(p / PToKGFactor)
}

// KGToP converts Kilograms to Pounds
func KGToP(kg Kilogram) Pound {
	return Pound(kg * PToKGFactor)
}
