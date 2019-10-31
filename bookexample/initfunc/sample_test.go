package main

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("Init function call")
}
func TestMethodA(t *testing.T) {
	MethodA()
}
