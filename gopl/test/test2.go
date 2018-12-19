package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var w io.Writer

func main() {
	w = os.Stdout
	f, ok := w.(*os.File)
	fmt.Println(f, ok)
	c, ok := w.(*bytes.Buffer)
	fmt.Println(c, ok)

	// rw := w.(io.ReadWriter)
	// w = new(ByteCounter)
	// rw = w.(io.ReadWriter)
}
