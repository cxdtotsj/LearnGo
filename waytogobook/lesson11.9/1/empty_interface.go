package main

import (
	"fmt"
)

var i = 5
var str = "ABC"

type person struct {
	name string
	age  int
}

type any interface{}

func main() {
	var val any
	val = 5
	fmt.Println("val value is :", val)

	val = str
	fmt.Println("val has value is :", val)

	pers1 := new(person)
	pers1.name = "Bear Chen"
	pers1.age = 55
	val = pers1
	fmt.Println("val has value is :", val)

	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string  %T\n", t)
	case bool:
		fmt.Printf("Type bool  %T\n", t)
	case *person:
		fmt.Printf("Type pointer to person  %T\n", t)
	default:
		fmt.Printf("Unexpected type  %T", t)
	}
}
