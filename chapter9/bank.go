package main

import (
	"fmt"
	"time"
)

// 一个函数在线性的程序中可以正常工作，如果在并发的情况下，
// 这个函数依然可以正确地工作的话，那么我们就说这个函数是并发安全的，
// 并发安全的函数不需要额外的同步工作。我们可以把这个概念概括为一些
// 方法和操作函数，如果这个类型是并发安全的话，那么所有它访问方法和
// 操作就都是并发安全

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

// balance变量被限制在了monitor gorountine中
func teller() {
	var balance int
	balance = 1
	fmt.Printf("Running ! %d\n", balance)
	for {
		select {
		case amount := <-deposits:
			balance += amount
		// 无阻塞的channel就是存的一瞬间必须拿！
		// 无阻塞的channel就是存的一瞬间必须拿！
		// 无阻塞的channel就是存的一瞬间必须拿！
		case balances <- balance:
			fmt.Printf("Has give %d", balance)
		}
	}
}

func init() {
	go teller()
}

func main() {
	done := make(chan struct{})

	time.Sleep(5 * time.Second)

	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		Deposit(100)
		done <- struct{}{}
	}()

	<-done
	<-done

	fmt.Printf("Hope is gonna be 300 %d\n", Balance())
}
