package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/hiboot/pkg/utils/crypto/base64"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// This sample use middleware to catch the request body and response body
//how to write a reverse proxy in go
//https://hackernoon.com/writing-a-reverse-proxy-in-just-one-line-with-go-c1edfa78c84b
func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

type DebugTransport struct {
}

func (DebugTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	b, err := httputil.DumpRequestOut(r, false)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(b))
	return http.DefaultTransport.RoundTrip(r)
}

// this middleware is used to proxy a https web server
func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")

		if r.Method != http.MethodGet {
			next.ServeHTTP(w, r)
			return
		}
		if strings.HasSuffix(r.URL.Path, "/v2/"){
			next.ServeHTTP(w, r)
			return
		}

		proxyHost := "10.160.210.111"
		username := "admin"
		password := "Harbor12345"
		repoName := "library/envoy"
		opType := "pull"
		token := RetrieveBearerToken(proxyHost, username, password, repoName, opType)
		fmt.Printf("Bearer token is %s\n", token.Token)

		proxyBaseUrl := &url.URL{
			Scheme: "https",
			Host:   proxyHost,
			Path:   "/",
		}
		fmt.Printf("request url before response: %#v\n", r.URL)
		fmt.Printf("proxy base url before response: %#v\n", proxyBaseUrl)
		proxy := httputil.NewSingleHostReverseProxy(proxyBaseUrl)

		fmt.Printf("request url path is %v\n", r.URL.Path)
		fmt.Printf("target url path is %v\n", proxyBaseUrl.Path)

		//r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		bearer := fmt.Sprintf("Bearer %s", token.Token)
		r.Header.Set("Authorization", bearer)
		r.Host = r.URL.Host

		fmt.Printf("The request url is %s\n", r.URL.String())

		config := GetLocalTLSConfig()
		proxy.Transport = &http.Transport{TLSClientConfig: config}

		proxy.ServeHTTP(w, r)
		fmt.Printf("request url after response: %#v\n", r.URL)

		log.Println("Executing middlewareOne again")
	})
}


func GetLocalTLSConfig() *tls.Config {
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
		InsecureSkipVerify: true,
		RootCAs:            rootCAs,
	}
	return config
}

type MyResponseWriter struct {
	http.ResponseWriter
	Buf *bytes.Buffer
}

func (mrw *MyResponseWriter) Write(p []byte) (int, error) {
	mrw.Buf.Write(p)
	return mrw.ResponseWriter.Write(p)
}


func Final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK\n"))
}

func main() {
	finalHandler := http.HandlerFunc(Final)

	//http.Handle("/", middlewareOne(middlewareTwo(finalHandler)))
	http.Handle("/", middlewareOne(finalHandler))
	http.ListenAndServe(":3000", nil)
}

func GetManifest(manifestUrl string, token *BearerToken) []byte {
	req2, err := http.NewRequest("GET", manifestUrl, nil)
	if err != nil {
		fmt.Println(err)
	}
	bt := fmt.Sprintf("Bearer %s", token.Token)
	fmt.Println(bt)
	req2.Header.Add("Authorization", bt)
	req2.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	client := getHttpClient()
	resp2, err := client.Do(req2)
	if err != nil {
		fmt.Println(err)
	}
	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		return []byte{}
	}
	return body2
}

func GetAuthorization(name, password string) string {
	return string(base64.Encode([]byte(fmt.Sprintf("%s:%s", name, password))))
}

func RetrieveBearerToken(hostname string, username string, password string, repoName string, opType string) *BearerToken {
	client := getHttpClient()
	urlString := fmt.Sprintf("https://%s/service/token?account=%s&scope=repository:%s:%s&service=harbor-registry", hostname, username, repoName, opType)
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/json")
	//req.Header.Add("Authorization", "YWRtaW46SGFyYm9yITIzNDU=") //admin:Harbor!2345
	req.Header.Add("Authorization", GetAuthorization(username, password)) //admin:Harbor!2345

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	token := &BearerToken{}
	json.Unmarshal(body, token)
	return token
}

func getHttpClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
	}
	return client
}

func ExtractDigest(body2 []byte) []string {
	fmt.Printf("The manifest content1 is %q\n", string(body2))
	digestList := make([]string, 0)
	rawIn := json.RawMessage(body2)
	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		panic(err)
	}
	m := &Manifest{}
	fmt.Printf("The manifest content2 is %q\n", bytes)
	err = json.Unmarshal(bytes, m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("manifest is : %#v", m)
	for _, l := range m.Layers {
		digestList = append(digestList, fmt.Sprintf("%v", l["digest"]))
	}
	return digestList
}

type BearerToken struct {
	Token       string `json:"token"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type Manifest struct {
	SchemaVersion int                      `json:"schemaVersion"`
	MediaType     string                   `json:"mediaType"`
	Config        Config                   `json:"config"`
	Layers        []map[string]interface{} `json:"layers"`
}
type Config struct {
	MediaType string `json:"mediaType"`
}
