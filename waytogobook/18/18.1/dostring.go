package main

import (
	"fmt"
)

func main() {
	// 如何修改字符串中的一个字符：
	str := "hello"
	c := []byte(str)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)

	// 如何获取字符串的子串：
	n, m := 0, 2
	substr := str[n:m]
	fmt.Println(substr)

	// 如何使用for或者for-range遍历一个字符串：
	for ix, ch := range str {
		fmt.Println(ix, "---->", ch)
	}

	// 拼接字符串
	str1 := "Hello "
	str2 := "World"
	str1 += str2
	fmt.Println(str1)
}
