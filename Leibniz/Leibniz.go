package main

import (
	"fmt"
	"math"
)

// 4*((-1)**n/(2n+1))

func main() {
	pi := 0.0

	for n := 0; n >= 0; n++ {
		func(n int64) {
			numerator := math.Pow(-1, float64(n))
			denominator := float64((2 * n) + 1)
			pi += numerator / denominator

		}(int64(n))
		fmt.Println(4 * pi)
	}

}
