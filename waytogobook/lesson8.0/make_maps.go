package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreate := make(map[string]float32)
	mapAssigned = mapLit

	mapCreate["key1"] = 4.5
	mapCreate["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Println(mapLit["one"])
	fmt.Println(mapCreate["key2"])
	fmt.Println(mapLit["two"])
	fmt.Println(mapLit["ten"])
}
