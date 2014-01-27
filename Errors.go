package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e *ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number : [ " + e + " ]"
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}

	z := f
	temp := z
	count := 0
	// for i := 0; i < 10; i++ {
	for {
		z = z - (math.Pow(z, 2)-f)/(2*z)
		if math.Abs(temp-z) < math.Pow(10, -14) {
			break
		}
		temp = z
		count++

		// fmt.Println(count)
		// fmt.Println(z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
