package main

import (
	"fmt"
)

// 是否重新分配内存取决于cap！！！
// 是否重新分配内存取决于cap！！！
// 是否重新分配内存取决于cap！！！

func remove(slice []int, i int) []int {
	// 这里可以看出，copy是不会开辟新的内存的，就看你给的source和dest数组是什么
	x1 := slice[i:]
	fmt.Printf("len is %d\n cap is %d\n slice is %v\n", len(x1), cap(x1), x1)
	x2 := slice[i+1:]
	fmt.Printf("len is %d\n cap is %d\n slice is %v\n", len(x2), cap(x2), x2)
	copy(slice[i:], slice[i+1:])
	fmt.Printf("len is %d\n cap is %d\n slice is %v\n", len(slice), cap(slice), slice)
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
