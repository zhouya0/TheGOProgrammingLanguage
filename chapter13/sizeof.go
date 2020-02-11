package main

import (
	"fmt"
	"unsafe"
)

var x struct {
	a bool
	b int16
	c []int
}

func main() {
	fmt.Println(unsafe.Sizeof(int16(0)))
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof([]int{1, 2, 3}))
	// unsafe.Alignof函数返回对应参数的类型需要对齐的倍数,也就是它的地址所对应的倍数
	// unsafe.Offsetof函数的参数必须是一个字段x.f，然后f字段相对于x起始地址的偏移量，包括可能的空洞
	// 对于x：
	// Size(x) = 32 Alignof(x) = 8
	// Sizeof(x.a) = 1 Alignof(x.a) = 1 Offsetof(x.a) = 0
	// Sizeof(x.b) = 2 Alignof(x.b) = 2 Offsetof(x.b) = 2
	// Sizeof(x.c) = 24 Alignof(x.c) = 8 Offsetof(x.c) = 8
	//  +---------*---------------------------+
	//  |  a |    |  b     |                  |
	//  +-------------------------------------+
	//  |            c(data)                  |
	//  +-------------------------------------+
	//  |            c(len)                   |
	//  +-------------------------------------+
	//  |            c(cap)                   |
	//  +-------------------------------------+
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b)
}
