package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	err := errors.New("abc")
	fmt.Println(err.Error())
	err2 := errors.New("efg")
	fmt.Println(err2.Error())
	fmt.Println(reflect.TypeOf(err))
	if err == err2 {
		fmt.Println(111)
	} else {
		fmt.Println(0)
	}
}
