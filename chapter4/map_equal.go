package main

import (
	"fmt"
)

func equalMap(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, v := range x {
		if yv, ok := y[k]; !ok || yv != v {
			return false
		}
	}
	return true
}

func main() {
	testMapx := map[string]int{"a": 1, "b": 2}
	testMapy := map[string]int{"a": 1, "b": 2}
	result := equalMap(testMapx, testMapy)
	fmt.Printf("Two value are %v equal\n", result)
}
