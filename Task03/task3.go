package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	chanel := make(chan int, 5)
	var res int
	for _, i := range array {
		wg.Add(1)
		go func(i int) {
			chanel <- i * i
			wg.Done()
		}(i)
	}
	for i := 0; i < 5; i++ {
		res += <-chanel
	}
	wg.Wait()
	fmt.Println(res)
}
