package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[string]string)
	m["zero"] = "sample"
	m["apple"] = "sample2"
	m["orange"] = "sample4"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println("key:", k, "value:", m[k])
	}
}
