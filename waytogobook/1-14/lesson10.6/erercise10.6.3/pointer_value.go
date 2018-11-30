package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) first() string {
	return p.firstName
}

func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}

func main() {
	p1 := Person{"a", "b"}
	fmt.Println(p1.first())
}
