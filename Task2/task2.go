package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for _, i := range array {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i * i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
