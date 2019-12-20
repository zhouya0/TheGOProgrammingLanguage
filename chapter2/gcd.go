package main

import (
	"fmt"
)

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	result := gcd(33, 45)
	fmt.Println(result)
}
