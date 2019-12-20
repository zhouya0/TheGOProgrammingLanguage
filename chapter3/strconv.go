package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	strA := fmt.Sprintf("%d", x)
	strB := strconv.Itoa(x)
	fmt.Println(strA, strB)
}

