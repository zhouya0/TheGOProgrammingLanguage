package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 断言一定是对interface进行断言
	// x.(T)类型断言有两种情况：
	// 第一种：
	// 1. 如果断言的类型T是一个具体类型
	// 2. 检查x的动态类型和T是否相同
	// 3. 如果相同，则返回x的动态值
	// 第二种：
	// 1. 如果断言的类型T是一个接口类型
	// 2. 检查x的动态类型是否满足T
	// 3. 如果满足，则还是返回x，但是x有类型T。
	type TestStruct struct {
		test string
	}
	var _ TestStruct

	var w1 io.Writer
	w1 = os.Stdout
	f := w1.(*os.File)
	fmt.Printf("Type is %T\n", f)

	var w2 io.Writer
	w2 = os.Stdout
	// w2.Read() will panic because io.Writer does not have this method
	r := w2.(io.ReadWriter)
	fmt.Printf("Type is %T\n", r)

}
