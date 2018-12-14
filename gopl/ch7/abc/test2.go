package main

import (
	"fmt"
	"sort"
)

// type a os.File

// type b io.Writer

// type c bytes.Buffer

// type d time.Duration

// var w io.Writer
// var rwc io.ReadWriteCloser

// var x interface{} = time.Now()

// func main() {
// 	w = os.Stdout
// 	w = new(bytes.Buffer)

// 	rwc = os.Stdout
// 	// rwc = new(bytes.Buffer)

// 	w = rwc
// 	// rwc = w
// 	fmt.Printf("%v", x)
// }

//sort 1. 序列的长度；2.两个元素的比较结果；3.交换两个元素的方式
type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	names := []string{"d", "c", "e", "a"}
	sort.Sort(StringSlice(names))
	fmt.Println(names)
}
