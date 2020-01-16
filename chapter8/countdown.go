package main

import (
	"fmt"
	"time"
)

func main() {
	counts := time.Tick(2 * time.Second)
	for n := 10; n > 0; n-- {
		fmt.Printf("Counting down to %d\n", n)
		<-counts
	}
}
