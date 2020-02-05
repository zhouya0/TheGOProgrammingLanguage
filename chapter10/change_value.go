package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type fooStruct struct {
	Name string
	Age  int
}

// rtype must be kept in sync with ../runtime/type.go:/^type._type.
// type rtype struct {
// 	size       uintptr
// 	ptrdata    uintptr  // number of bytes in the type that can contain pointers
// 	hash       uint32   // hash of type; avoids computation in hash tables
// 	tflag      tflag    // extra type information flags
// 	align      uint8    // alignment of variable with this type
// 	fieldAlign uint8    // alignment of struct field with this type
// 	kind       uint8    // enumeration for C
// 	alg        *typeAlg // algorithm table
// 	gcdata     *byte    // garbage collection data
// 	str        nameOff  // string form
// 	ptrToThis  typeOff  // type for pointer to this type, may be zero
// }

func core(i interface{}) {
	// 搞懂	e := (*emptyInterface)(unsafe.Pointer(&i))
	p := unsafe.Pointer(&i)
	f := (*fooStruct)(p)
	fmt.Printf("struct type is %T\n", f)
	// 下面这个一定会报错！！！
	// fmt.Println(f.Name)
	fmt.Printf("type is %T\n", p)
	fmt.Println(p)
}

func main() {
	// 一个变量就是一个可寻址的内存空间，里面存储了一个值，并且
	// 存储的值可以通过内存地址来进行更新
	x := 2
	a := reflect.ValueOf(2)
	fmt.Printf("type is %T\n", a)
	fmt.Println(a)
	b := reflect.ValueOf(x)
	fmt.Printf("type is %T\n", b)
	c := reflect.ValueOf(&x)
	fmt.Printf("type is %T\n", c)
	d := c.Elem()
	fmt.Printf("type is %T\n", d)
	// 代码中完全就是靠flag标志位来确定是否可以Addr,
	fmt.Println(a.CanAddr())

	test := 1
	core(test)
}
