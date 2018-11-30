package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputfile := "./product.txt"
	outputfile := "./product_copy.txt"
	buf, err := ioutil.ReadFile(inputfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File error: %s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputfile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}
}
