package main

import (
	"fmt"
)

func main() {
	testSlice := make(map[string]int)
	testSlice["one"] = 1
	fmt.Println(testSlice)
	copyMap := testSlice
	copyMap["two"] = 2
	fmt.Println(testSlice)

	x := make([]int, 3)
	x = []int{1, 2, 3}
	x = append(x, 5) // append就厉害了哦，它是会发生一个拷贝的。
	x = append(x, 6)
	fmt.Println(x)
	y := x
	y[1] = 3
	fmt.Println(y)
	fmt.Println(x)
	var z []int = make([]int, 5)
	copy(z, x)
	z[1] = 999
	fmt.Println(x)
}
