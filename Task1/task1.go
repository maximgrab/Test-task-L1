package main

import "fmt"

type Human struct {
	name string
}

func (human Human) PrintName() {
	fmt.Printf("Name: %v.\n", human.name)
}

type Action struct {
	Human
}

func main() {
	action := Action{Human{name: "Abcde"}}
	action.PrintName()
	action.Human.PrintName()

}
