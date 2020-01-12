package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	naturals := make(chan int)
	suqares := make(chan int)
	go func() {
		for {
			naturals <- 1
		}
	}()

	go func() {
		for {
			n := <-naturals
			s := n * n
			suqares <- s
		}
	}()

	go func() {
		fmt.Println(<-suqares)
	}()
	wg.Wait()
}
