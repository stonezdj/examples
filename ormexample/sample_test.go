package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestMyFunc(t *testing.T) {
	userInfo := url.UserPassword("stonezdj", "@zhu88=jie")
	fmt.Println("Sample!")
	myurl := url.URL{
		Scheme: "postgres",
		Host:   "10.0.0.1",
		User:   userInfo,
	}

	fmt.Println(myurl.String())
}
