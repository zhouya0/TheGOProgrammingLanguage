package main

import (
	"strconv"
	"strings"
	"sync"
)

type Set struct {
	data map[int]interface{}
	// without this, it will have
	// fatal error: concurrent map writes
	mu sync.RWMutex
}

func New() Set {
	data := make(map[int]interface{})
	return Set{data: data}
}

func (s *Set) Add(num int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[num] = nil

}

func (s *Set) Delete(num int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, num)
}

func (s *Set) String() string{
	result := ""
	for k,_ := range s.data {
		result = strings.Join([]string{result, strconv.Itoa(k)}, "")
	}
	return result

}

func main() {
	test := New()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(){
		for i:=0; i<1000;i++ {
			test.Add(i)
		}
		wg.Done()
	}()
	go func(){
		for i:=0; i<1000;i++ {
			test.Add(i)
		}
		wg.Done()
	}()
	wg.Wait()
}
