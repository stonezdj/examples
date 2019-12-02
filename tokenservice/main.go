package main

import (
	"fmt"
)

func main() {

	hostname := "10.160.210.111"
	username := "admin"
	repoName := "library/envoy"
	opType := "push,pull"
	password := "Harbor12345"

	token := RetrieveBearerToken(hostname, username, password, repoName, opType)

	fmt.Printf("Token is:%v\n", token.Token)

	//use bearer token to request the manifest
	manifestUrl := fmt.Sprintf("https://%s/v2/%s/manifests/latest", hostname, repoName)

	body2 := GetManifest(manifestUrl, token)
	digestList := ExtractDigest(body2)

	for _, d := range digestList {
		fmt.Println(d)
	}

}
