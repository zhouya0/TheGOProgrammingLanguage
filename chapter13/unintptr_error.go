package main

import (
	"unsafe"
)

var x struct {
	a bool
	b int
	c []int
}

type test struct{}

func main() {
	// 下面的代码是及其不安全的
	// 当一个变量被移动，所有的保存改变变量旧地址的指针必须同时被更新为变量移动后的地址
	// 但是uintptr类型的临时变量只是一个普通数字，所以其值不应该被改变
	// 最后导致`unsafe.Pointer(tmp)`得到的东西可能并不是x.b(x被移动)
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb := (*int)(unsafe.Pointer(tmp))
	*pb = 42

	// 下面的代码也是不安全的，这里并没有指针引用new创建新的变量，因此改语句执行完之后，垃圾收集器有权马上回收其内存
	// 空间，所以返回的pT将是无效的地址。
	pT := uintptr(unsafe.Pointer(new(test)))
	if pT != 0 {
	}
}
