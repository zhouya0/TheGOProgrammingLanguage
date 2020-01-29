package main

import (
	"fmt"
)

func main() {
	for {
		go func() {
			fmt.Print("1")
		}()

		fmt.Print("0")
	}
}
