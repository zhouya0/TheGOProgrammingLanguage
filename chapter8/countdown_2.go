package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return abort.")
	select {
	// select的核心逻辑就是，满足任何一个case就通过否则就一直阻塞！
	// 一直阻塞！！！
	// 一直阻塞！！！
	case <-time.After(10 * time.Second):
		// Do nothing means stop.
	case <-abort:
		fmt.Println("Aborting!!!")
		return
	default:

	}

	launch()
}

func launch() {
	fmt.Println("Launching!!!")
}
