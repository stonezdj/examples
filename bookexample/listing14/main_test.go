package listing14

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestTemp(t *testing.T) {

	cust := Customer{
		ContactInfo{
			"Beijing",
			"13641072766",
			"100065"},
		"Daojun"}
	fmt.Printf("%+v", cust)
}
func TestTemp2(t *testing.T) {
	cust := Customer{
		ContactInfo: ContactInfo{
			Address:  "Beijing",
			Mobile:   "13641072766",
			Postcode: "100065"},
		Name: "Daojun"}
	fmt.Printf("%+v", cust)
}

func TestMethod3(t *testing.T) {
	var array1 [5]string
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	array1 = array2
	array2[1] = "Orange"
	for _, item := range array1 {
		fmt.Printf(item + "\t")
	}

	fmt.Print("\n")
	for _, item := range array2 {
		fmt.Printf(item + "\t")
	}
}

func TestMethod4(t *testing.T) {
	s1 := []int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1)

	s2 := s1[1:]
	fmt.Println(len(s2), cap(s2), s2)

	for i := range s2 {
		s2[i] += 20
	}

	fmt.Println(s1)
	fmt.Println(s2)

	s2 = append(s2, 4)
	for i := range s2 {
		s2[i] += 10
	}

	fmt.Println(s1)
	fmt.Println(s2)
}
func TestMethod5(t *testing.T) {
	s1 := [5]int{1, 2, 3}
	fmt.Println(len(s1), cap(s1), s1)
}

func TestMethod7(t *testing.T) {
	configMap := map[string]interface{}{
		"key1": "123",
		"key2": "sample",
	}

	fmt.Printf("result : %+v\n", reflect.TypeOf(configMap["key1"]))

	if _, ok := configMap["key1"].(float64); ok {
		fmt.Println("It is int")
	} else {
		fmt.Println("Not a int")
	}

}

func TestMethod9(t *testing.T) {
	s1 := []byte(`{"key1": 123, "key2": "sample"}`)
	var maps map[string]interface{}

	json.Unmarshal(s1, &maps)

	for key, value := range maps {
		fmt.Printf("key=%v, value=%v\n", key, value)
	}
}
func TestMethod10(t *testing.T) {
	var jsonBlob = []byte(`[
		{"Name":"Playtypus", "Order": "Monotremata"},
		{"Name":"Quoll", "Order":"Dasyuromorphia"}
	]
	`)

	type Animal struct {
		Name    string
		Order   string
		Comment string
	}

	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}

func TestMethodAccessTag(t *testing.T) {
	type User struct {
		Name  string `mytag:"MyName"`
		Email string `mytag:"MyEmail"`
	}
	u := User{"Bob", "bob@mycompany.com"}
	r := reflect.TypeOf(u)

	for _, fieldName := range []string{"Name", "Email"} {
		field, found := r.FieldByName(fieldName)
		if !found {
			continue
		}
		fmt.Printf("\tField: User.%s\n", fieldName)
		fmt.Printf("\tWhole tag value : %q\n", field.Tag)
		fmt.Printf("\tValue of 'mytag': %q\n", field.Tag.Get("mytag"))
	}
}
