package main

import (
	"fmt"
	"time"
)

func sleep(delay time.Duration) {
	startTime := time.Now()
	for {
		if time.Now().Sub(startTime) > delay {
			return
		}
	}
}
func main() {
	sleep(time.Second * 5)
	fmt.Println("10")
}
