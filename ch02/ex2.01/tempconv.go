package tempconv

import "fmt"

// Celsius Temperature °C
type Celsius float64

// Fahrenheit Temperature °F
type Fahrenheit float64

// Kelvin Temperature °K
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

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
