package main

import (
	"fmt"
)

func main() {
	str := "Chinese : 日本语"
	fmt.Printf("The length of str is : %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("字符顺序 %d 是 : %c\n", pos, char)
	}
}

func Season(month int) (seasonName string) {
	switch month {
	case 1, 2, 3:
		return "spring"
	case 4, 5, 6:
		return "summer"
	case 7, 8, 9:
		return "autumn"
	case 10, 11, 12:
		return "winter"
	}
	return "Season unknow"
}
