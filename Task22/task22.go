package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(2 << 21)
	b := big.NewInt(2 << 21)
	var res big.Int
	fmt.Println(res.Quo(a, b))
	fmt.Println(res.Sub(a, b))
	fmt.Println(res.Add(a, b))
	fmt.Println(res.Mul(a, b))
}
