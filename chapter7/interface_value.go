package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	// fmt的“%T”操作可以获取接口的动态类型
	// 接口的动态类型！！！
	// 接口的动态类型！！！
	// 而它的实现，其实也就是使用了反射
	fmt.Printf("%T\n", w)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// w.Write([]byte("Hello"))
	w = os.Stdout
	fmt.Printf("%T\n", w)
	w.Write([]byte("Hello"))
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)
}
