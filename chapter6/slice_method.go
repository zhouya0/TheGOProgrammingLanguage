package main

import (
	"fmt"
)

type SliceTest []int

func (s SliceTest) Sum() int {
	var r int
	for _, v := range s {
		r = r + v
	}
	return r
}

func main() {
	t := SliceTest{1, 2, 3, 4}
	fmt.Println(t.Sum())
}
