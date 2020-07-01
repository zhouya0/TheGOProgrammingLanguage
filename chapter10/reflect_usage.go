package main

import (
	"fmt"
	"reflect"
)

// 相关TypeOf和ValueOf的最基本的东西。
//package main
//
//import (
//	"fmt"
//	"reflect"
//)
//
//func main() {
//	var num float64 = 1.2345
//
//	fmt.Println("type: ", reflect.TypeOf(num))
//	fmt.Println("value: ", reflect.ValueOf(num))
//}
//
//运行结果:
//type:  float64
//value:  1.2345


// 从Field函数调用来讲，ValueOf返回的总是一个Value，但是TypeOf返回的就是一个StructuredField。

// 动态调用函数

type T struct {}

func (t *T) Foo(){
	fmt.Println("Foo")
}

func (t *T) Bar() {
	fmt.Println("Bar")
}

// struct tag解析
type S struct {
	A int `json:"aaa" test:"testaaa"`
	B string `json:"bbb" test:"testbbb"`
}

func main() {
	// 动态处理函数
	t := &T{}
	reflect.ValueOf(t).MethodByName("Foo").Call(nil)
	reflect.ValueOf(t).MethodByName("Bar").Call(nil)

	// struct tag解析
	s := S{
		A: 123,
		B: "hello",
	}
	tt := reflect.TypeOf(s)
	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)
		if json, ok := field.Tag.Lookup("json"); ok {
			fmt.Println(json)
		}
		if test, ok := field.Tag.Lookup("test"); ok {
			fmt.Println(test)
		}
	}

	// 通过kind处理不同的分支
	a := 1
	k := reflect.TypeOf(a)
	switch k.Kind() {
	case reflect.Int:
		fmt.Println("I'm int")
	case reflect.String:
		fmt.Println("I'm string")
	}
}