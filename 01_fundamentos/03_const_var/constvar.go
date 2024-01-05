package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.14159
	var raio = 3.3

	// Forma reduzida
	// area := PI * raio * raio
	area := PI * math.Pow(raio, 2)
  fmt.Printf("A área é %f", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = a
		d = b
	)

	fmt.Println(c, d)
	var e, f bool = true, false
	fmt.Println(e, f)
	g, h, i := 2, false, "epa"
	fmt.Println(g, h, i)
}
