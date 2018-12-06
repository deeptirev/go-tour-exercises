package main

import (
	"fmt"
)

func Squarert(x float64) float64 {
	z := 1.0
	for {
		if (abs(z*z-x) < 1e-15) {
			return z;
		}
		z -= (z*z - x) / (2 * z)
	}
}

func abs(x float64) float64 {
	if (x < 0) {
		return -(x)
	} else {
		return x
	}
}

func main() {
	fmt.Println(Squarert(16))
}
