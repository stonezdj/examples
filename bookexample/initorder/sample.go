package main

import "fmt"

func init() {
	fmt.Println("Called first in order of declaration")
	initCounter++
	fmt.Printf("Init counter %v\n", initCounter)
}

var initCounter = 32

// func MethodA() {
// 	fmt.Println("Call method A")
// }

func init() {
	fmt.Println("Called second in order of declaration")
	initCounter++
}

func main() {
	fmt.Println("Does nothing of any significance")
	fmt.Printf("Init Counter: %d\n", initCounter)
	MethodA()
}
