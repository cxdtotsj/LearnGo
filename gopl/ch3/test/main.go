package main

import (
	"fmt"
)

// func main() {
// 	// 定义数组类型
// 	q := [...]int{1, 2, 3}
// 	fmt.Printf("%T\n", q)
// }

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 //英镑
	RMB                 // 人民币
)

func main() {
	// symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	// r := [...]int{5: -2, -1}
	// fmt.Println(RMB, symbol[RMB])
	// fmt.Println(len(symbol))
	// fmt.Println(len(r))
	// fmt.Printf("%T\n", r)
	// fmt.Printf("%v\n", r)

	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%T\n", a)
}
