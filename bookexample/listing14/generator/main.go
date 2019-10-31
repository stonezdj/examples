package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	c := boring("boring!")
// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("You say: %q\n", <-c) //Receive expression is just a value
// 	}
// 	fmt.Println("You're boring; I'am leaving.")
// }

// func main() {
// 	joe := boring("Joe")
// 	ann := boring("Ann")
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(<-joe)
// 		fmt.Println(<-ann)
// 	}
// 	fmt.Println("You're both boring; I'am leaving.")
// }

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'am leaving.") 
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) //Expression to be send can be any suitable value
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
