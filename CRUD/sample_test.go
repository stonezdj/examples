package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestFormatString(t *testing.T) {
	//cfgMap := GetMapFromRequest(`{"token_expiration":107374182400000}`)

	cfgMap := GetMapFromRequest(`{"token_expiration":32}`)
	fmt.Println(GetStrValueOfAnyType(cfgMap["token_expiration"]))
}

func GetStrValueOfAnyType(value interface{}) string {
	var strVal string
	if _, ok := value.(map[string]interface{}); ok {
		b, err := json.Marshal(value)
		if err != nil {
			fmt.Printf("can not marshal json object, error %v", err)
			return ""
		}
		strVal = string(b)
	} else {
		switch value.(type) {
		case float64:
			fmt.Println("Got float64")
			strVal = fmt.Sprintf("%.0f", value)
		case float32:
			fmt.Println("Got float32")
			strVal = fmt.Sprintf("%.0f", value)
		default:
			fmt.Println("Got Default")
			strVal = fmt.Sprintf("%v", value)
		}

	}
	return strVal
}

func GetMapFromRequest(request string) map[string]interface{} {
	var cfgMap map[string]interface{}
	json.Unmarshal([]byte(request), &cfgMap)
	return cfgMap
}
