//15.4

package main

import (
	"bufio"
	"fmt"
	"os"
)

type page struct {
	Title string
	Body  []byte
}

func (p *page) save() {
	outputFile, outputError := os.OpenFile("book.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("open error")
		return
	}

	outputWriter := bufio.NewWriter(outputFile)
	for i := range p.Body {
		outputWriter.WriteString(p.Title)
		outputWriter.WriteByte(p.Body[i])
	}

}

func main() {
	p := new(page)
	p.Title = "test"
	p.Body = [1,2,3,4]

}
