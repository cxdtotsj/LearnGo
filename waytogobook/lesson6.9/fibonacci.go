// 练习 6.10

package main

import (
	"fmt"
	"strings"
)

func main() {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	fmt.Println(addBmp("file"))
	fmt.Println(addJpeg("file1"))
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
