package main

import "fmt"

type twoInts struct {
	a int
	b int
}

func main() {
	two1 := new(twoInts)
	// two1 := twoInts{10, 20}	// two1 是 结构体类型 也可以
	two1.a = 10
	two1.b = 20
	fmt.Printf("The sum is : %d \n", two1.addThem())
	fmt.Printf("the sum of addThemTwo is : %d\n", two1.addThemTwo(30))

	two2 := twoInts{100, 200}
	fmt.Println(two2.addThem())
	fmt.Println(two2.addThemTwo(300))
}

func (tn *twoInts) addThem() int {
	return tn.a + tn.b
}

func (tn *twoInts) addThemTwo(param int) int {
	return tn.a + tn.b + param
}
