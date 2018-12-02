/*
1.接口 IDance 实现了两个方法
2.函数 DuckDance 使用接口类型 IDance 作为参数
3.定义 类型 Bird
4.只有实现了 IDace 定义方法的类型，才能作为 DuckDance 函数的参数

*/

package main

import (
	"fmt"
)

type IDance interface {
	Quick()
	Walk()
}

func DuckDance(duck IDance) {
	for i := 0; i <= 3; i++ {
		duck.Quick()
		duck.Walk()
	}
}

type Bird struct {
}

func (b *Bird) Quick() {
	fmt.Println("I'm quicking!")
}

func (b *Bird) Walk() {
	fmt.Println("I'm walking!")
}

func main() {
	b := new(Bird)
	DuckDance(b)
}
