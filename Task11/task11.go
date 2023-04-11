package main

import "fmt"

func main() {
	stringSequence1 := []string{"cat", "cat", "dog", "cat", "tree", "sheep"}
	stringSequence2 := []string{"cow", "cat", "bush", "dog", "house"}
	stringSet1 := make(map[string]bool)
	for _, i := range stringSequence1 {
		stringSet1[i] = true
	}
	stringSet2 := make(map[string]bool)
	for _, i := range stringSequence2 {
		stringSet2[i] = true
	}
	setsIntersection := make(map[string]bool)
	for i := range stringSet1 {
		if stringSet2[i] {
			setsIntersection[i] = true
		}
	}
	for i := range setsIntersection {
		fmt.Println(i)
	}
}
