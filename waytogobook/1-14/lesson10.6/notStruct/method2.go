package main

import (
	"fmt"
)

type intVector []int

func (v intVector) sum() (s int) {
	for _, x := range v {
		s += x
	}
	return s
}

func main() {
	// intv := intVector{1, 2, 3, 4, 5}
	intv := make(intVector, 3)
	intv[0] = 1
	intv[1] = 2
	intv[2] = 3
	s := intv.sum()
	fmt.Println(s)
}
