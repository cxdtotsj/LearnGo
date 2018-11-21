package main

import (
	"fmt"
)

type shaper interface {
	Area() float32
}

type square struct {
	side float32
}

func (sq *square) Area() float32 {
	return sq.side * sq.side
}

type rectangle struct {
	length, width float32
}

func (rct rectangle) Area() float32 { //rct is struct value
	return rct.length * rct.width
}

func main() {

	// the one sample
	// sq := new(square)
	// sq.side = 0.5
	// sq := &square{0.6}
	// var sh shaper
	// sh = sq
	// fmt.Println(sh.Area())

	// the two sample
	r := rectangle{0.3, 0.5}
	q := &square{0.8}
	shapes := []shaper{r, q}
	for n := range shapes {
		fmt.Println("shape details : ", shapes[n])
		fmt.Println("Area of shape is : ", shapes[n].Area())
	}
}
