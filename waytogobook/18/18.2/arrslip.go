package main

func main() {
	// 创建
	// arr1 := new([len]type)
	// slice1 := make([]type,len)

	//初始化
	// arr1 := [...]type{i1,i2,i3,i4,i5}
	// arrKeyValue := [len]type{i1:val1,i2:val2}
	// var slice1 []type = arr1[start:end]

	// 如何截断数组或者切片的最后一个元素：
	// line = line[:len(line)-1]

	// 如何使用for或者for-range遍历一个数组（或者切片）：
	// for i :=0;i<len(arr);i++{
	// 	... = arr[i]
	// }
	// for ix,value := range arr {
	// 	...
	// }

	//如何在一个二维数组或者切片arr2Dim中查找一个指定值V：
	// found := false
	// Found: for row:=range arr2Dim{
	// 	for colum := range arr2Dim[row] {
	// 		if arr2Dim[row][colum] == v {
	// 			found = true
	// 			break Found
	// 		}
	// 	}
	// }
}
