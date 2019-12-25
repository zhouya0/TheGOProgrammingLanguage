package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonemptyNew(strings []string) []string {
	newSlice := make([]string, 0)
	for _, v := range strings {
		if v != "" {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func main() {
	testSlice := []string{"one", "", "three", "four", "five"}
	noneSlice := nonempty(testSlice)
	fmt.Println(noneSlice)
	testSliceNew := []string{"one", "", "three", "four", "five", "", "seven"}
	result := nonemptyNew(testSliceNew)
	fmt.Println(result)
	fmt.Printf("The length of the new array is %d\n", len(result))
}
