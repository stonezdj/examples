package main

import (
	"bytes"
	"net/http"
)

// example middleware to peek the content of request body.
// https://www.alexedwards.net/blog/making-and-using-middleware
func enforceXMLHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for a request body
		if r.ContentLength == 0 {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		// Check its MIME type
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		if http.DetectContentType(buf.Bytes()) != "text/xml; charset=utf-8" {
			http.Error(w, http.StatusText(415), 415)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	finalHandler := http.HandlerFunc(final)

	http.Handle("/", enforceXMLHandler(finalHandler))
	http.ListenAndServe(":3000", nil)
}

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK\n"))
}
