package main

import (
	"fmt"
	"testing"
)

func TestSample(t *testing.T) {
	for i := 1; i < 100; i++ {
		fmt.Println(i)
	}
}
