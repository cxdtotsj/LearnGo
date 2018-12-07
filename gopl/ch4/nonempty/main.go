package main

import (
	"fmt"
)

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	// fmt.Printf("%q\n", data)

	// 更新data
	data = nonempty(data)
	fmt.Printf("%q\n", data)
}

// func nonempty(strings []string) []string {
// 	i := 0
// 	for _, s := range strings {
// 		if s != "" {
// 			strings[i] = s
// 			i++
// 		}
// 	}
// 	return strings[:i]
// }

func nonempty(strings []string) []string {
	out := strings[:0]
	for _, r := range strings {
		if r != "" {
			out = append(out, r)
		}
	}
	return out
}
