package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	inputChanel := make(chan int, 5)
	outputChanel := make(chan int, 5)
	wg.Add(1)
	go func() {
		defer close(inputChanel)
		for _, i := range array {
			inputChanel <- i
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		defer close(outputChanel)
		for i := range inputChanel {
			outputChanel <- i * i
		}
		wg.Done()
	}()
	for i := range outputChanel {
		fmt.Println(i)
	}
	wg.Wait()
}
