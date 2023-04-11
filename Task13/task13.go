package main

import "fmt"

func main() {
	a := 5
	b := 8
	//Способ 1
	fmt.Printf("\n%v %v\n", a, b)
	a += b
	b = a - b
	a -= b
	fmt.Printf("%v %v\n", a, b)
	//Способ 2
	fmt.Printf("\n%v %v\n", a, b)
	a, b = b, a
	fmt.Printf("%v %v\n", a, b)
	//Способ 3
	fmt.Printf("\n%v %v\n", a, b)
	a ^= b
	b ^= a
	a ^= b
	fmt.Printf("%v %v\n", a, b)
}
