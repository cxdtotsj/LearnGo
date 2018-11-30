package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)

	time.Sleep(1e9)

}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			fmt.Printf("%v", open)
			break
		}
		fmt.Printf("%s\n", input)
	}
}
