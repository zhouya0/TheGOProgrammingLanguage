package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex
	balance int
	wg sync.WaitGroup
)

// 一般来讲，被mutex所保护的变量是在mutex变量声明之后立即声明的。
// Mutex锁只能申请一个！！！
// Mutex锁只能申请一个！！！
// Mutex锁只能申请一个！！！


// 这里使用channel会发生阻塞，因为一个goroutine在请求锁，另一个goroutine在等待你完成。
func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

func main() {

	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(200)
		fmt.Println("=", Balance())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(100)
	}()

	wg.Wait()
	fmt.Printf("Hope is gonna be 300 %d\n", Balance())
}
