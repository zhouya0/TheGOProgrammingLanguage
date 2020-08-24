package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Set struct {
	data map[int]interface{}
	mu sync.RWMutex
}

func New() Set {
	data := make(map[int]interface{})
	return Set{data: data}
}

func (s *Set) Add(num int) {
	s.data[num] = nil
}

func (s *Set) Delete(num int) {
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
	test.Add(3)
	test.Add(4)
	test.Add(2)
	test.Add(3)
	fmt.Printf("The string is: %v\n",  test.String())

	testString := string([]byte{65, 66})
	fmt.Println(testString)
	testString = string(65)
	fmt.Println(testString)


}
