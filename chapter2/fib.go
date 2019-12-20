package main

import "fmt"

func fib(n int) int {
	if n == 1 {
		return 0
	}
	x := 0
	y := 1
	for n != 2 {
		x, y = y, x+y
		n--
	}
	return y
}

func main() {
	test := fib(7)
	fmt.Println(test)
}
