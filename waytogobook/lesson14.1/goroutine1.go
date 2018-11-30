package main

import (
	"fmt"
	"time"
)

func main() {

}

func longwait() {
	fmt.Println("Beginning longwait()")
	time.Sleep(5 * 1e9)
	fmt.Println("End of lognwait()")
}

func shortwait() {
	fmt.Println("Beginning shortwait()")
}
