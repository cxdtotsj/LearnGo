//创建
// type struct1 struct{
// 	field1 type1
// 	field2 type2
// 	...
// }

// ms := new(struct1)

// 初始化：
// ms := &struct1{10,15.5,"Chris"}

// 使用构建函数初始化结构体
package main

import (
	"fmt"
)

type struct1 struct {
	n    int
	f    float32
	name string
}

func main() {
	ms := newStruct1(1, 1.1, "Chris")
	fmt.Println(ms)
}

func newStruct1(n int, f float32, name string) *struct1 {
	return &struct1{n, f, name}
}
