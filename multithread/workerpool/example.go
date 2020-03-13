package main

import (
	"fmt"
	"time"
	"sync"
)
var wg sync.WaitGroup
var exit = make(chan bool)
func main() {


	reqs := make(chan *Request, 100)

	for i:=0; i<100; i=i+1 {
		r := &Request{
			Username: fmt.Sprintf("stonezdj%d", i),
			Address:  fmt.Sprintf("Beijing%d", i),
		}
		fmt.Printf("Adding work %v\n", i)
		wg.Add(1)
		reqs <- r

	}

	go Serve(reqs, exit)
	wg.Wait()

	fmt.Println("All complete")
}

type Request struct {
	Username string
	Address string
}

const MaxOutStanding = 5
var sem = make(chan int, MaxOutStanding)

func handle(clientRequests chan *Request){
	for req := range clientRequests{
		r := req
		sem <- 1
		process(r)
		<- sem
	}
	exit <- true


}

func process(r *Request){
	time.Sleep(1*time.Second)
	fmt.Printf("Username %v, address %v\n", r.Username, r.Address)
	wg.Done()
}

func Serve(clientRequests chan *Request, exit chan bool){
	for i := 0; i < MaxOutStanding; i++ {
		go handle(clientRequests)
	}
	<-exit
}