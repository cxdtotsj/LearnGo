package main

import (
	"fmt"
)

func main() {
	// a := [...]int{1, 2, 3, 4, 5}
	// reverse(a[:])
	// a := []int{1, 2, 3, 4, 5}
	a := [5]int{1, 2, 3, 4, 5} // array
	reverse(&a)
	fmt.Println(a)
}

// func reverse(s []int) {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// }

// exercise 4.3
func reverse(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
