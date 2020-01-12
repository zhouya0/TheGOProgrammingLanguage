package main

import (
	"fmt"
	"sync"
	"time"
)

func spinner(delay time.Duration, wg *sync.WaitGroup) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
	// the Done() for sync.WaitGroup always stay in goroutine
	defer wg.Done()
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go spinner(100*time.Millisecond, &wg)
	wg.Wait()

	// const n = 45
	// fibN := fib(n)
	// fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
