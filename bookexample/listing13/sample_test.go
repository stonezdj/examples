package main

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestRegulate(t *testing.T) {
	fmt.Println(math.Mod(-1, 3))
}

func TestMethodInitialize(t *testing.T) {
	fmt.Println("Testing message")
	var sample []int64
	var pointSample *[]int64
	fmt.Printf("Sample :%v\n", sample)
	fmt.Printf("pointer sample :%v\n", pointSample)
}

func TestIntegerLiteral(t *testing.T) {
	a := 0600
	b := 40
	fmt.Printf("Value: %d\n", (a + b))
	c := 072.40
	fmt.Printf("Value c=%f\n", c)
}

func TestMethod(t *testing.T) {
	sample := []string{"a", "b", "c"}
	newSample := make([]string, 0)
	for _, item := range sample {
		item = "'" + item + "'"
		newSample = append(newSample, item)
	}
	fmt.Printf("join result:%v", strings.Join(newSample, ","))

}
