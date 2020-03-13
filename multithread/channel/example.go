package main

import (
	"fmt"
	"time"
)

/*
 Example to create a channel to receive signal and sync goroutine
*/

func main(){
	c := make(chan int)
	go func(){
		sort()
		c <- 1
	}()

	doSomethingForAWhile()
	fmt.Println("Example")
	<-c
	fmt.Println("sort finished")
}

func sort(){
	fmt.Println("Start to sort")
	time.Sleep(15*time.Second)
	fmt.Println("End to sort")

}

func doSomethingForAWhile(){
	fmt.Println("Sleep 10 seconds begin")
	time.Sleep(10*time.Second)
	fmt.Println("Sleep 10 seconds end")

}