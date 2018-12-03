package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	copyfile("target,txt", "source.txt")
	fmt.Printf("Copt done!")
}

func copyfile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
