package main

import (
	"fmt"
	// unsafe是用于指针类型的转换
	"unsafe"
)

type iface struct {
	itab, data uintptr
}

func main() {
	var a interface{} = nil

	var b interface{} = (*int)(nil)

	x := 5
	var c interface{} = (*int)(&x)

	// 类型转化！！！
	// ia 是真正意义上的全空，动态类型和动态值都为空
	// ib 是有静态类型的，但是没有动态值，所以itab不为空，但是data为空
	// ic 是有静态类型和动态值的，所以两个都不为空
	// *iface是一个指针！！！
	// *iface是一个指针！！！
	// *iface是一个指针！！！
	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	ic := *(*iface)(unsafe.Pointer(&c))

	fmt.Println(ia, ib, ic)

	fmt.Println(*(*int)(unsafe.Pointer(ic.data)))

	// s := Teststring{"test"}
	// test := Testint(s)
	// fmt.Println(test)
}

// 不用想了 这种类型转换肯定不行
// type Teststring struct {
// 	dataString string
// }

// func (t Teststring) String() {
// 	fmt.Println(t.dataString)
// }

// type Testint struct {
// 	data int
// }

// func (t Testint) String() {
// 	fmt.Println(t.data)
// }
