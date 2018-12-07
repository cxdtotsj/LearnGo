package main

import (
	"fmt"
)

// func main() {
// 	// 定义数组类型
// 	q := [...]int{1, 2, 3}
// 	fmt.Printf("%T\n", q)
// }

// type Currency int

// const (
// 	USD Currency = iota // 美元
// 	EUR                 // 欧元
// 	GBP                 //英镑
// 	RMB                 // 人民币
// )

// func main() {
// 	// symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
// 	// r := [...]int{5: -2, -1}
// 	// fmt.Println(RMB, symbol[RMB])
// 	// fmt.Println(len(symbol))
// 	// fmt.Println(len(r))
// 	// fmt.Printf("%T\n", r)
// 	// fmt.Printf("%v\n", r)

// 	a := [...]int{0, 1, 2, 3, 4, 5}
// 	fmt.Printf("%T\n", a)
// }

// 比较两个 slice(非bytes类型) 是否相等
// func equal(x, y []string) bool {
// 	if len(x) != len(y) {
// 		return false
// 	}
// 	for i := range x {
// 		if x[i] != y[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

//内置append
// var runes []rune

// func main() {
// 	for _, r := range "Hello, 世界" {
// 		runes = append(runes, r)
// 	}
// 	fmt.Printf("%q\n", runes)
// }

// 对 key的类型为 string 的map 进行排序

// func main() {

// 	ages := map[string]int{"bob": 32, "alent": 50, "chen": 28}

// 	// var names []string
// 	names := make([]string, 0, len(ages)) //优化性能
// 	for name := range ages {
// 		names = append(names, name)
// 	}
// 	sort.Strings(names)
// 	for _, name := range names {
// 		fmt.Printf("%s\t%d\n", name, ages[name])
// 	}
// }

// 判断两个 map 中的 k,v 是否相等

func main() {
	x := map[string]int{"A": 0}
	y := map[string]int{"A": 32}
	fmt.Printf("%t\n", equal(x, y))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
