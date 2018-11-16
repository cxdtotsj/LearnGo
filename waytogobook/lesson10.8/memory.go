package main

import (
	"fmt"
	"runtime"
)

var m runtime.MemStats

func main() {
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
