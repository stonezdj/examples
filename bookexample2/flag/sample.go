package main

import (
	"flag"
	"fmt"
)

var heapsize int

func init() {
	flag.IntVar(&heapsize, "heapsize", 1234, "help message for flagname")
}

func main() {
	flag.Parse()
	fmt.Printf("heapsize:=%v\n", heapsize)
}
