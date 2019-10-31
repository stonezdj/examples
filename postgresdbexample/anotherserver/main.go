package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://10.160.218.91:9999", nil)
	req.Close = true

	//req.Header.Add("Accept-Encoding", "gzip")

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// content, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println("done!")
}
