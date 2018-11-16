package main

import (
	"fmt"
	"strconv"
)

type TwoTnts struct {
	a int
	b int
}

func main() {
	// two1 := new(TwoTnts)
	// two1.a = 12
	// two1.b = 22
	two1 := &TwoTnts{100, 200}
	fmt.Printf("two1 is : %v\n", two1)
	fmt.Println("two1 is :", two1)
	fmt.Printf("two1 is : %T\n", two1)
	fmt.Printf("two1 is : %#v\n", two1)
}

func (tn *TwoTnts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
