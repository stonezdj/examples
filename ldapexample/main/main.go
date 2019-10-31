package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"gopkg.in/ldap.v2"
)

func main() {

	var RootCACert string
	//RootCACert = "/Users/daojunz/Documents/vol/ldapcert2/ca.crt /Users/daojunz/Documents/vol/ldapcert2/ldap.crt"
	RootCACert = "/Users/daojunz/Documents/vol/ldapcert2/ca.crt"
	//RootCACert = "/Users/daojunz/Documents/vol/ldapcert2/notary-signer-ca.crt"
	var certPool *x509.CertPool

	certPool = x509.NewCertPool()
	for _, caCertFile := range strings.Split(RootCACert, " ") {
		if pem, err := ioutil.ReadFile(caCertFile); err != nil {
			fmt.Println("cert read fail!")
		} else {
			if !certPool.AppendCertsFromPEM(pem) {
				fmt.Println("Cert append faile!")
			}
		}
	}

	//Connect
	//l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "10.4.140.203", 389))
	l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", "10.4.140.203", 636), &tls.Config{ServerName: "10.4.140.203", InsecureSkipVerify: false})
	//l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", "10.4.140.203", 636), &tls.Config{ServerName: "ldap-service", InsecureSkipVerify: false, RootCAs: certPool})
	if err != nil {
		fmt.Println("Failed on dialTLS")
		log.Fatal(err)
	} else {
		fmt.Println("Connect success")
	}

	defer l.Close()

	//Bind DN
	err = l.Bind("cn=admin,dc=example,dc=org", "admin")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Bind success!")
	}

	//Search inetOrgPerson
	searchRequest := ldap.NewSearchRequest(
		"dc=example,dc=org",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(structuralObjectClass=inetOrgPerson))",
		[]string{"dn", "cn"},
		nil,
	)

	sr, err := l.Search(searchRequest)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Search result:")
	for _, entry := range sr.Entries {
		fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}

}
