package listing14

import (
	"fmt"
)

// ContactInfo ..
type ContactInfo struct {
	Address  string
	Mobile   string
	Postcode string
}

// Customer ..
type Customer struct {
	ContactInfo
	Name string
}

// Main ..
func Main() {
	fmt.Println("This is the hello-world!")
	cust := Customer{
		ContactInfo{
			"Beijing",
			"13641072766",
			"100065",
		},
		"Daojun",
	}

	fmt.Println("%+v", cust)
}
