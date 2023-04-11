package main

import "fmt"

func main() {
	tempSequence := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempSet := make(map[int][]float32)
	for _, i := range tempSequence {
		tempSet[int(i/10)*10] = append(tempSet[int(i/10)*10], i)
	}
	fmt.Println(tempSet)
}
