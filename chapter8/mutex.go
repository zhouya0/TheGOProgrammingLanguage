package main

import (
	"fmt"
	"sync"
)

type MutexInt struct {
	myMap int
	mu sync.Mutex
}

func (m *MutexInt) Inc(){
	// 如果不要这个lock，那么这个结果将是不稳定的。
	m.mu.Lock()
	defer m.mu.Unlock()
	m.myMap++
}

func main() {
	myChan := make(chan string, 100)
	// 可以看到，这里的sync.Mutex是不需要实例化的。
	myInt := MutexInt{myMap:1}
	for i:=0; i<100; i++ {
		go func (){
			myInt.Inc()
			myChan <- ""
		}()

	}
	for i:=0 ; i<100; i++ {
		<-myChan
	}
	fmt.Println(myInt.myMap)
}
