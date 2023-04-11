package main

import "fmt"

func main() {
	stringSequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)
	for _, i := range stringSequence {
		set[i] = true
	}
	for i := range set {
		fmt.Println(i)
	}
}
