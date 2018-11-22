package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading,exiting program")
		return
	}

	// fmt.Printf("Your name is %s", input)
	// //For Unix:test with delimiter "\n",for windows: test with "\r\n"
	// switch input {
	// case "Philip\r\n":
	// 	fmt.Println("Welcome Philip!")
	// case "Chris\r\n":
	// 	fmt.Println("Welcome Chris")
	// case "Ivo\r\n":
	// 	fmt.Println("Welcome Ivo!")
	// default:
	// 	fmt.Println("You are not welcome here! Goodbye!")

	// }

	//version 2
	switch input {
	case "Philip\r\n":
		fallthrough
	case "Ivo\r\n":
		fallthrough
	case "Chris\r\n":
		fmt.Printf("Welocme %s\n", input)
	default:
		fmt.Println("You are not welcome here! Goodbye!")
	}
}
