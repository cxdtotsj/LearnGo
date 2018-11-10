//斐波那锲计算函数
// 练习6.5

package main

import "fmt"

func main() {
	result := 0
	n := 0
	for i := 10; i >= 0; i-- {
		result, n = loca(i)
		fmt.Println("The result is :", result, "The location is", n)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func loca(n int) (res, i int) {
	res, i = fibonacci(n), n
	return
}
