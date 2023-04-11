package main

import "fmt"

func typePrint(a interface{}) {
	fmt.Printf("%T\n", a)
}

func main() {
	typePrint(53)
}
