// exercise of the book:The way to go

package main

import (
	"LearnGo/waytogobook/exercise"
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println(exercise.Add(42, 13))
	exercise.GetVersion()
	fmt.Println(exercise.Mult) //exercise 包的常量

	// 4.5
	var goos string = runtime.GOOS
	fmt.Printf("The System is : %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("PATH is\n %s\n", path)
}
