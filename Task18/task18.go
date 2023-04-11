package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	i     int
	mutex sync.Mutex
}

func (counter *Counter) inc() {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.i += 1
}
func main() {
	var wg sync.WaitGroup
	var counter Counter
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				counter.inc()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter.i)
}
