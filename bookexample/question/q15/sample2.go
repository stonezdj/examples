package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	Map map[int]int
}

func newSafeMap(size int) *SafeMap {
	sm := new(SafeMap)
	sm.Map = make(map[int]int)
	return sm
}

func (sm *SafeMap) readMap(key int) int {
	sm.RLock()
	defer sm.RUnlock()
	value := sm.Map[key]
	fmt.Printf("value=%v\n", value)
	return value
}

func (sm *SafeMap) writeMap(key int, value int) {
	sm.Lock()
	defer sm.Unlock()
	sm.Map[key] = value
}

func main() {
	safeMap := newSafeMap(10)

	for i := 0; i < 100000; i++ {
		go safeMap.writeMap(i, i)
		go safeMap.readMap(i)
	}

}
