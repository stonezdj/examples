package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
	}

	//request for bearer token
	urlString := "https://10.160.215.19/service/token?account=admin&scope=repository:library/envoy:push,pull&service=harbor-registry"
	//u, err := url.Parse(urlString)
	//if err != nil {
	//	log.Error(err)
	//}
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "YWRtaW46SGFyYm9yITIzNDU=")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	token := &BearerToken{}
	json.Unmarshal(body, token)
	fmt.Printf("Token is:%v\n", token.Token)

	//use bearer token to request the manifest
	manifestUrl := "https://10.160.215.19/v2/library/envoy/manifests/latest"

	req2, err := http.NewRequest("GET", manifestUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	bt := fmt.Sprintf("Bearer %s", token.Token)
	fmt.Println(bt)
	req2.Header.Add("Authorization", bt)
	req2.Header.Add("Accept", "application/vnd.docker.distribution.manifest.v2+json")
	resp2, err := client.Do(req2)

	if err != nil {
		fmt.Println(err)
	}

	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The manifest content is %q", string(body2))

}

type BearerToken struct {
	Token       string `json:"token"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
