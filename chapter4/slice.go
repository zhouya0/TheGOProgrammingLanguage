package main

import (
	"fmt"
)

func printSlice(s []string) {
	fmt.Printf("len of slice is %d\n", len(s))
	fmt.Printf("cap of slice is %d\n", cap(s))
}

func main() {
	testSlice := []string{"zero", "one", "two", "three", "four", "five", "six", "seven"}
	printSlice(testSlice)
	cuttedSlice := testSlice[2:5]
	// The cuttedSlice will have capcity of 6 because it uses testSlice's memory
	printSlice(cuttedSlice)
	cuttedSlice[0] = "test"
	fmt.Println(cuttedSlice)
	fmt.Println(testSlice)
}
