package main

import (
	"fmt"
)

func main() {
	var testint uint8 = 4
	// 永远记住咯，它的逻辑是：把左边的东西移右边的位数。
	fmt.Printf("Result is %08b\n", testint<<1)
	fmt.Printf("Result is %08b\n", testint>>1)
	fmt.Printf("Result is %08b\n", 1<<testint)
	fmt.Printf("Result is %08b\n", 1>>testint)
}

