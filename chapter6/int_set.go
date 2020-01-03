package main

import (
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	fmt.Printf("word is %d, bit is %d \n", word, bit)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	fmt.Printf("add %d , word is %d, bit is %d \n", x, word, bit)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	// 这里的“|”是或操作，其意义是，把某一位没有被置“1”的给置一了。这样每一个1就代表着这个数存在。
	s.words[word] |= 1 << bit
	fmt.Println(s.words)
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func main() {
	var test IntSet
	test.Add(5)
	test.Add(6)
	test.Add(7)
	// 所以增加63这个操作，就会产生一次超大的运算。
	test.Add(63)
	test.Add(65)
	test.Add(78)
	fmt.Println(test.Has(65))
}
