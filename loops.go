package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	
	for ; math.Abs((z*z - x) / (2*z)) > 0.000001;  z -= (z*z - x) / (2*z) {
		fmt.Println(z)
	}
	
	return z
}

func main() {
	fmt.Println(Sqrt(6))
}
