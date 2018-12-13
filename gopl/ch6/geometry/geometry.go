package main

import (
	"fmt"
	"math"
)

type point struct{ x, y float64 }
type path []point

func main() {
	p := point{1, 2}
	q := point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	//path
	perim := path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1}}
	fmt.Println(perim.Distance())
}

func Distance(p, q point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p point) Distance(q point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (path path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
