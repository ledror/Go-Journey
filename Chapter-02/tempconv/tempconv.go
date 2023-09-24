package tempconv

import "fmt"

type (
	Celcius    float64
	Fahrenheit float64
	Kelvin     float64
)

const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
	AbsoluteZeroK Kelvin  = 0
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
)

func (c Celcius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
