package main

import (
	"fmt"

	"LearnGo/waytogobook/lesson10.3/struct_pack"
)

func main() {
	// point
	struct1 := new(structPack.Expstruct)
	struct1.Mi1 = 10
	struct1.Mi2 = 16.2

	fmt.Printf("M1 = %d \n", struct1.Mi1)
	fmt.Printf("M2 = %.1f \n", struct1.Mi2)

	//value type
	var struct2 structPack.Expstruct
	struct2.Mi1 = 20
	struct2.Mi2 = 26.2

	fmt.Printf("M1 = %d \n", struct2.Mi1)
	fmt.Printf("M2 = %.1f \n", struct2.Mi2)
}
