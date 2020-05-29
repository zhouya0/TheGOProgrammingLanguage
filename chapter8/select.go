package main

import (
	"fmt"
)


// 来源于https://github.com/kubernetes/kubernetes/pull/79777

// 原本的函数是这样的：
// 	catchError := func(err error) {
//		lock.Lock()
//		defer lock.Unlock()
//		if firstError == nil {
//			firstError = err
//		}
//	}
// 但是我们可以用select来解决


type mych struct {
	ch chan string
}


// 有data就输入，输入不了就离开
func (m *mych) sendData (s string) {
	select{
	case m.ch <- s:
	default:
	}
}

// 有data就拿，无data就离开
func (m *mych) getData () string {
	select{
	case s := <- m.ch:
		return s
	default:
		return ""
	}
}

func newMych() mych{
	return mych{make(chan string, 1)}
}


func main() {
	m := newMych()
	tests := []string{"test1", "test2", "test3"}
	for _,v := range tests{
		m.sendData(v)
	}
	fmt.Println(m.getData())
}
