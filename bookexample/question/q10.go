package main

import "fmt"

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	if stu == nil {
		fmt.Println("It is nil")
	}
	return stu
}

func main() {
	//	if live() == (*Student)(nil) {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
