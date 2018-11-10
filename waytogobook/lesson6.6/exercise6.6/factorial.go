package main

import "fmt"

func main() {
	result := 0
	for i := 0; i <= 3; i++ {
		result = factorial(i)
		fmt.Println("The factorial is ", result)
	}
}

func factorial(n int) (res int) {
	if n == 0 {
		res = 1
	} else {
		res = factorial(n) * factorial(n-1)
	}
	return
}
