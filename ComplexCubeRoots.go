package main

import "fmt"
import "math"
import "math/cmplx"

func Cbrt(x complex128) complex128 {
	// z := complex128(x + 1i)
	z := x
	temp := z
	count := 0
	// for i := 0; i < 10; i++ {
	for {
		z = z - (cmplx.Pow(z, 3)-x)/(3*cmplx.Pow(z, 2))
		if cmplx.Abs(temp-z) < math.Pow(10, -14) {
			break
		}
		temp = z
		count++

		// fmt.Println(count)
		// fmt.Println(z)
	}
	fmt.Println(count)

	return z
}

func main() {
	fmt.Println(Cbrt(200 + 101i))
}
