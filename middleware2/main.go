package main

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// This sample use middleware to catch the request body and response body

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
		config := &tls.Config{
			InsecureSkipVerify: true,
		}

		proxy.Transport = &http.Transport{TLSClientConfig: config}
		proxy.ServeHTTP(w, r)
		//destConn, err := net.DialTimeout("tcp", "www.apache.org:80", 10*time.Second)
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusServiceUnavailable)
		//	return
		//}
		//w.WriteHeader(http.StatusOK)
		//
		//hj, ok:=w.(http.Hijacker)
		//if !ok {
		//	http.Error(w, "web server doesn't support hijacking", http.StatusInternalServerError)
		//	return
		//}
		//
		//clientConn, _, err := hj.Hijack()
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		//
		//go transfer(destConn, clientConn)
		//go transfer(clientConn, destConn)

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
