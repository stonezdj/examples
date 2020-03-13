package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	fmt.Fprintf(w, "This is an example server.\n")
	io.WriteString(w, "This is an example server.\n")
}
// put the server.crt and server.key in the current directory.
func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":9443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
