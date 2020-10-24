package main

import (
	"fmt"
	"reflect"
)

func main2() {

	a := int(222)
	b := int32(11)
	c := float32(55.01)
	fmt.Println("Welcome to Integers!")
	fmt.Printf("a : type is %T\n", a)
	fmt.Printf("b : type is %v\n", reflect.TypeOf(b))
	fmt.Printf("c : type is %T\n", c)

}
