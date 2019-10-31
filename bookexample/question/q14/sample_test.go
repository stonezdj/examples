package sample

import (
	"fmt"
	"testing"
)

func TestMethod(t *testing.T) {
	fmt.Println("Hello world!")
	test := Student{
		internalValue: "sample1",
		ExternalValue: "next sample",
	}

	fmt.Printf("message need to print,%+v\n", test)
}
