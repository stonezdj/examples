package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	operation string
}

func returnsError() error {
	var p error = nil
	if bad() {
		p = errors.New("Failed on bad")
	}
	return p
}

func bad() bool {
	return false
}

func main() {
	p := returnsError()
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}
