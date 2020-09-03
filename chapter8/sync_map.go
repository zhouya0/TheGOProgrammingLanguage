package main

import (
	"sync"
)


func main() {
	syncMap := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(){
		do(syncMap)
		wg.Done()
	}()
	go func(){
		do(syncMap)
		wg.Done()
	}()
	wg.Wait()

}

func do (m sync.Map) {
	i := 0
	for i < 10000 {
		m.Store(1,1)
		i ++
	}
}