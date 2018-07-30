// Package tempconv performs Celsius and Fahrenheit temperature computations
package main

// New types are only "named types," in that their underlying type
// is an existing type
type Celsius float64
type Fahrenheight float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheight { return Fahrenheight(c*9/5 + 32) }
func FToC(f Fahrenheight) Celsius { return Celsius((f - 32) * 5 / 9) }
