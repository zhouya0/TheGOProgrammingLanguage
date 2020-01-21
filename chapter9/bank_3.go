package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex
	balance int
)

// 一般来讲，被mutex所保护的变量是在mutex变量声明之后立即声明的。

// 这里为什么会发生死锁呢？需要研究！！！
func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := Balance()
	mu.Unlock()
	return b
}

func main() {
	done := make(chan struct{})

	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	fmt.Printf("Hope is gonna be 300 %d\n", Balance())
}
