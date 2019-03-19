package main

import (
	"fmt"
	"math"
)

func fact(n float64) float64 {

	var fact float64 = 1
	var i float64
	for i = 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}

func main() {

	var k float64
	var pi float64 = 0
	for k = 0; k >= 0; k++ {
		//num := new(big.Int).SetUint64(uint64(math.Pow(2,math.Float64frombits(k))*math.Pow(float64(fact(k)),2)))
		//div := new(big.Int).SetUint64(fact(2*k + 1))
		//pi := new(big.Int).Div(num, div)
		num := math.Pow(2, float64(k)) * math.Pow(float64(fact(k)), 2)
		div := float64(fact(2*k + 1))
		pi += num / div
		fmt.Println(2 * pi)
	}
}
