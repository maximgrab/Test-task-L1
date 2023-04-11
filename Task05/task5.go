package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	N = 10
)

func main() {
	timer := time.NewTimer(time.Second * N)
	ticker := time.NewTicker(time.Second)
	valChanel := make(chan int)
	stopChanel := make(chan int)
	go func() {
		for {
			select {
			case <-stopChanel:
				return
			case val := <-valChanel:
				fmt.Println(val)
			}
		}
	}()
	for {
		select {
		case <-ticker.C:
			valChanel <- rand.Int()
		case <-timer.C:
			stopChanel <- 1
			return
		}
	}
}
