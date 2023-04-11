package main

import (
	"fmt"
	"sync"
)

type SyncedMap struct {
	mutex sync.RWMutex
	imap  map[int]int
}

func (m *SyncedMap) Set(key int, val int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.imap[key] = val
}
func (m *SyncedMap) Get(key int) (int, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	val, ok := m.imap[key]
	return val, ok
}
func main() {
	var wg sync.WaitGroup
	syncedMap := &SyncedMap{
		imap: make(map[int]int),
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			syncedMap.Set(i, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(syncedMap.Get(i))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
