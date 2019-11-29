package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// This sample use middleware to catch the request body and response body
//how to write a reverse proxy in go
//https://hackernoon.com/writing-a-reverse-proxy-in-just-one-line-with-go-c1edfa78c84b
func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

// this middleware is used to proxy a https web server
func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		url, err := url.Parse("https://10.160.215.19/harbor/projects/1/repositories")
		if err != nil {
			log.Println(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(url)
		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = url.Host

		rootCAs, _ := x509.SystemCertPool()
		if rootCAs == nil {
			rootCAs = x509.NewCertPool()
		}

		localCertFile := "/Users/daojunz/Downloads/cert/ca.crt"
		certs, err := ioutil.ReadFile(localCertFile)
		if err != nil {
			log.Fatalf("Failed to append %q to RootCAs: %v", localCertFile, err)
		}

		if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
			log.Println("No certs appended, using system certs only")
		}

		config := &tls.Config{
			InsecureSkipVerify: false,
			RootCAs:            rootCAs,
		}

		proxy.Transport = &http.Transport{TLSClientConfig: config}
		proxy.ServeHTTP(w, r)

		//next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

type MyResponseWriter struct {
	http.ResponseWriter
	Buf *bytes.Buffer
}

func (mrw *MyResponseWriter) Write(p []byte) (int, error) {
	mrw.Buf.Write(p)
	return mrw.ResponseWriter.Write(p)
}

// this middleware is used to peek the request content and response content.
func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		if r.URL.Path != "/" {
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		} else {
			log.Printf("the body is %v", string(body))
		}

		mrw := &MyResponseWriter{
			ResponseWriter: w,
			Buf:            &bytes.Buffer{},
		}
		next.ServeHTTP(mrw, r)

		log.Printf("The response is %v\n", mrw.Buf.String())

		log.Println("Executing middlewareTwo again")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK123\n"))
}

func main() {
	finalHandler := http.HandlerFunc(final)

	http.Handle("/", middlewareOne(middlewareTwo(finalHandler)))
	http.ListenAndServe(":3000", nil)
}
