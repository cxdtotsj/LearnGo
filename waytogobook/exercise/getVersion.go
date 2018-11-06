package exercise

import (
	"fmt"
	"runtime"
)

func GetVersion() {
	fmt.Printf("%s", runtime.Version())
	fmt.Println(Mult)
}
