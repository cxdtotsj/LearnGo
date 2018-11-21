package main

import "fmt"

type Element interface{}

type vector struct {
	a []Element
}

func (p *vector) At(i int) Element {
	return p.a[i]
}

func (p *vector) Set(i int, e Element) {
	p.a[i] = e
}

func main() {
	v := new(vector)
	fmt.Println(*v)
	v.Set(1, "abc")
	fmt.Println(v)
}
