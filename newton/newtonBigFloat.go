package main

import (
	"fmt"
	"math/big"
)

func Pow(a uint64, e *big.Float) *big.Float {
	result := Zero().Copy(e)
	//one := big.NewFloat(1)
	//eOne := big.NewFloat(0)
	//eOne = eOne.Sub(a,one)
	i := big.NewFloat(0)
	for i.Cmp(a-1) < 0 {
		result = Zero().Mul(result, a)
	}
	return result
}

func Zero() *big.Float {
	r := big.NewFloat(0.0)
	r.SetPrec(256)
	return r
}

func Mul(a, b *big.Float) *big.Float {
	return Zero().Mul(a, b)
}

func factBIG(n float64) float64 {

	var fact float64 = 1
	var i float64
	for i = 1; i <= n; i++ {
		fact = fact * i
	}
	return fact
}

func main() {

	k := big.NewFloat(2)
	//pi := big.NewFloat(0)
	//
	//num := big.NewFloat(0)
	//den := big.NewFloat(0)
	//
	//for {
	//	num = Pow(2, k)

	//}
	fmt.Println(Pow(2, k))

	//for k = 0; k >= 0; k++ {
	//	num := math.Pow(2, float64(k))*math.Pow(float64(factBIG(k)),2)
	//	div := float64(factBIG(2*k + 1))
	//	pi += num/div
	//	fmt.Println(2*pi)
	//}
}
