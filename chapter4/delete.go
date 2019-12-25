package main

import (
	"fmt"
)

func remove(slice []int, i int) []int {
	// 这里可以看出，copy是不会开辟新的内存的，就看你给的source和dest数组是什么
	copy(slice[i:], slice[i+1:])
	slice = slice[:len(slice)-1]
	return slice
}

// 更简单的方式，不需要调用copy，但是会改变原数组的顺序
func removeRaw(slice []int, i int) []int {
	slice[i] = slice[len(slice)]
	slice = slice[:len(slice)-1]
	return slice
}

func main() {
	testSlice := []int{1, 2, 3, 4, 5, 6, 7}
	result := remove(testSlice, 2)
	fmt.Println(result)
	fmt.Println(testSlice)
}
