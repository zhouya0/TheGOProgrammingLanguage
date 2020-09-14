package main

import "fmt"

type Stack struct {
	list []int
}

func (s *Stack) Pop() int{
	length := len(s.list)
	if length == 0 {
		return 0
	}
	r := s.list[length-1]
	s.list = s.list[:length-1]
	return r
}

func (s *Stack) Push(i int) {
	s.list = append(s.list, i)
}


type Queue struct {
	stack1 Stack
	Stack2 Stack
}

func (q *Queue) inQueue(i int) {
	q.stack1.Push(i)
}

func (q *Queue) outQueue() int{
	if len(q.Stack2.list) != 0 {
		return q.Stack2.Pop()
	} else {
		for len(q.stack1.list) != 0 {
			q.Stack2.Push(q.stack1.Pop())
		}
		return q.Stack2.Pop()
	}
}

func main() {
	testQueue := Queue{}
	testQueue.inQueue(3)
	testQueue.inQueue(4)
	testQueue.inQueue(5)
	fmt.Println(testQueue.outQueue())
	fmt.Println(testQueue.outQueue())
	fmt.Println(testQueue.outQueue())
}