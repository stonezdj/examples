// +build prod

package main

const (
	MONGODB_HOST            = "mongodb://prod.vmware.com:27017"
	MONGODB_DATABASE        = "productdb"
	MONGODB_CONNECTION_POOL = 5
	API_PORT                = 8080
)
