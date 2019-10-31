package main

import "fmt"

type I interface {
	walk()
}

type A struct {
	name string
}

func (a A) walk() {}

type B struct {
	name string
}

func (b B) walk() {}

func main() {
	var s interface{} = nil
	var p *string = nil
	if q, ok := s.(*string); ok {
		_ = q
		fmt.Println("casted")
	} else {
		fmt.Println("cast failed")
	}

	_ = p

}
