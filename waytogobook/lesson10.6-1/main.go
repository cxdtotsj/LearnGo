package main

import (
	"fmt"

	"LearnGo/waytogobook/lesson10.6-1/person"
)

func main() {
	p := new(person.Person)
	p.SetFirstName("Eric")
	fmt.Println(p.FirstName())
}
