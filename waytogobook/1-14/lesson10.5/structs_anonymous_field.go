package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    // anonymous
	innerS // anonymous
}

func main() {
	//指针
	out1 := new(outerS)
	out1.b = 1
	out1.c = 2.3
	out1.int = 3
	out1.in1 = 4
	out1.in2 = 5
	fmt.Println(out1)

	// 字面量结构体指针
	out2 := &outerS{11, 12.5, 13, innerS{21, 22}}
	fmt.Println(out2)

	// 字面量结构体
	out3 := outerS{100, 100.5, 110, innerS{200, 201}}
	fmt.Println(out3)
}
