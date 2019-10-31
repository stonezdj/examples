package main

import (
	"net/http"

	"github.com/fromkeith/gorest"
)

type HelloService struct {
	gorest.RestService `root:"/tutorial/"`
	helloWorld         gorest.EndPoint `method:"GET" path:"/hello-world/" output:"string"`
}

func (h HelloService) HelloWorld() string {
	return "Hello world"
}

func main() {
	gorest.RegisterService(new(HelloService))
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(":8787", nil)
	//visit http://localhost:8787/tutorial/hello-world/
}
