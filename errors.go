package main

import (
	"fmt"
	"math"
)

/* func Sqrt(x float64) (float64, error) {
	return 0, nil
} */

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	
	if (x < 0) {
		return 0.0, ErrNegativeSqrt(x)
	}
	
	z := 1.0
	
	for ; math.Abs((z*z - x) / (2*z)) > 0.000001;  z -= (z*z - x) / (2*z) {
		//fmt.Println(z)
	}
	
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
