// exercise of the book:The way to go

package main

import (
	"LearnGo/waytogobook/exercise/exercise"
	"fmt"
)

func main() {
	sum, mult, diff := exercise.Gethasvar(5, 8)
	fmt.Println(sum, mult, diff)
}
