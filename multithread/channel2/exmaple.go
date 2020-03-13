package main

import (
	"fmt"
	"time"
)

/**
	Use time out to sync goroutine
 */
func main() {

	done := make(chan bool)

	go func() {
		longTimeMethod(done)
	}()

	select {
		case <- done:
			fmt.Println("Complete longTimeMethod")
		case <- time.After(20*time.Second):
			fmt.Println("We can not wait for so long, exit firstly")
	}


}


func longTimeMethod(done chan bool){
	fmt.Println("Start to sleep 15 seconds")
	time.Sleep(15*time.Second)
	done <- true
	fmt.Println("End to sleep 15 seconds")
}