package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for math.Abs((z*z - x) / (2*z)) > x / 1000000 {
		z -= (z*z - x) / (2*z)
		fmt.Println("yuuupi")
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
