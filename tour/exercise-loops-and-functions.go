package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	var last float64
	for {
		last = z
		z -= (z*z - x) / (2 * z)
		fmt.Println("z = %j", z)
		if math.Abs(z-last) <= 0.00000001 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println("result is ", Sqrt(9))
}
