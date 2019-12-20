package main

import (
	"fmt"
)

func main() {
	var testint uint8 = 4
	fmt.Printf("Result is %08b\n", testint<<1)
	fmt.Printf("Result is %08b\n", testint>>1)
}