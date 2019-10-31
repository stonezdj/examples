package main

import (
	"fmt"
	"sync"
)

type singleton struct {
}

func (s *singleton) CallMethod() {
	fmt.Println("This is something to print")
}

var instance *singleton
var once sync.Once

// GetInstance ...
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	singleton := GetInstance()
	singleton.CallMethod()

}
