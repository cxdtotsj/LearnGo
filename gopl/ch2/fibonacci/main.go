package main

import (
	"fmt"
)

func main() {
	x, y := 0, 1
	n := 10
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Println(x)
	}
}

// 计算第 N 个数的值
func fibonacci(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
