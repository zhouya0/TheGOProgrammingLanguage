package main

// 为什么要使用链表，因为是链地址法啊，所以肯定是
// 哈希数组上每个都是一个链表啊！！！
// 哈希数组上每个都是一个链表啊！！！
// 哈希数组上每个都是一个链表啊！！！
type MpNode struct {
	Data Dict
	Next *MpNode
}

type Dict struct {
	Key string
	Value interface{}
}

func newNodeHead() *MpNode {
	node := new(MpNode)
	node.Data.Key = "头key"
	node.Data.Value = "头value"
	node.Next = nil
	return node
}

func (node *MpNode) data(k string, v interface{}) *MpNode {
	if node == nil {
		node = newNodeHead()
	}
	node.Data.Key = k
	node.Data.Value = v
	return node
}

func (node *MpNode) getByKey(k string) interface{} {
	if node.Next == nil {
		return nil
	}
	for node.Next != nil {
		if node.Next.Data.Key == k {
			return node.Next.Data.Value
		} else {
			node = node.Next
		}
	}
	return nil
}

func (node *MpNode) add(k string, v interface{}) {
	if node.getByKey(k) != nil {
		return
	}
	for node.Next != nil {
		node = node.Next
	}
	node.Next = node.Next.data(k,v)
}

func (node *MpNode) Length() int {
	if node == nil {
		return 0
	}
	i := 0
	for node.Next != nil {
		i++
		node = node.Next
	}
	return i
}

var Arr [16]*MpNode


func NewHash() {
	for i := 0; i < 16; i ++{
		Arr[i] = newNodeHead()
	}
}

func SetKey(k string, v interface{}) {
	num := hashNum(k)
	Arr[num].add(k,v)
}

func GetKey(k string) interface{} {
	num := hashNum(k)
	return Arr[num].getByKey(k)
}

func hashNum(key string) int {
	var index int = 0
	index = int(key[0])

	for k := 0; k < len(key); k++ {
		index *= (1103515245 + int(key[k]))
	}
	index >>= 27
	index &= 16-1
	return index
}
