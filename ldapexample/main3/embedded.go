package main

import (
	"fmt"
)

// User ...
type User struct {
	Username string
	Address  string
}

// StudentUser ...
type StudentUser struct {
	User
}

func main() {
	stu := StudentUser{
		User{
			Username: "daojun",
			Address:  "Beijing",
		},
	}
	fmt.Printf("this is the studentuser: %+v", stu)
	fmt.Printf("user:%+v", stu.User)

}
