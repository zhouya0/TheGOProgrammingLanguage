package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	naturals := make(chan int)
	suqares := make(chan int)
	go func() {
		for x := 0; ; x++ {
			naturals <- x
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
		for {
			fmt.Println(<-suqares)
			time.Sleep(1000 * time.Millisecond)
		}
	}()
	wg.Wait()
}
