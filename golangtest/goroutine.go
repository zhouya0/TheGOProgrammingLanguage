package main

import (
	"sync"
	"fmt"
)

func main() {
	ch := make(chan int, 10)
	for i:=0; i<5; i++ {
		ch <- i
	}
	close(ch)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// for range ch 这样的遍历，是没有v,ok := range ch这样的写法的
		for j:=0; j<20; j++{
			data, ok := <- ch
			fmt.Println(data)
			fmt.Println(ok)
		}
	}()
	wg.Wait()
	fmt.Println("over")
}

