package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	url string
)

func init() {
	flag.StringVar(&url, "url", "http://www.baidu.com", "Which URL do we want to parse?")
	flag.Parse()
}

func main() {
	err := send(url)
	if err != nil {
		panic(err)
	}
}

func send(link string) error {
	client := http.Client{}
	response, err := client.Get(link)
	if err != nil {
		return err
	}
	if response == nil {
		return err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}
