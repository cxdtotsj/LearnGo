package main

import "fmt"

type employee struct {
	salary float32
}

func (em *employee) giveRaise(per float32) float32 {
	return em.salary * (1.0 + per)
}

func main() {
	emp1 := employee{3000.0}
	fmt.Println(emp1.giveRaise(0.8))

	emp2 := &employee{18000.0}
	fmt.Println(emp2.giveRaise(0.8))
}
