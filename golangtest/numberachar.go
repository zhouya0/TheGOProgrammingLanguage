package main

import (
	"fmt"
	"sync"
)

var testNum  = []int{1,2,3,4,5,6,6,7,87,8,9,9,0,0}
var testChar = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N"}


func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(){
		for _,v := range testNum {
			// 如果这里的数字的长度大于char的长度，那么就会发生阻塞，那么程序就是无法完成的了。
			ch <- 1
			fmt.Print(v)

		}
		wg.Done()
	}()
	go func(){
		for _,v := range testChar {
			<- ch
			fmt.Print(v)
		}
		wg.Done()
	}()


	wg.Wait()
}