package main

import (
	"fmt"
	"sync"
	"time"
)

type UserAges struct {
	sync.Mutex
	ages map[string]int
}

// func NewUserAges() *UserAges {
// 	return &UserAges{
// 		ages: map[string]int{},
// 	}
// }

func (ua *UserAges) Set(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	time.Sleep(100 * time.Second)
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	ages := UserAges{ages: map[string]int{}}
	ages.Set("Tom", 99)
	go func() {
		ages.Set("Tom", 1000)
	}()
	fmt.Printf("get age %v", ages.Get("Tom"))
}
