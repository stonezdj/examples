package main

import (
	"fmt"
	"reflect"
)

type MyErr string

func (m MyErr) Error() string {
	return "big fail"
}

func doSomething(i int) error {
	switch i {
	default:
		return nil
	case 1:
		var p *MyErr
		return p
	case 2:
		return (*MyErr)(nil)
	case 3:
		var err error
		return err
	case 4:
		var p *MyErr
		return error(p)
	case 5:
		return (error)(nil)
	}
}

func main() {
	for i := 0; i <= 5; i++ {
		err := doSomething(i)
		fmt.Println(i, err, err == nil)
	}

	var x float64 = 3.4
	fmt.Println("value:", reflect.ValueOf(x).String())
}
