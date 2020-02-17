package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	Name string
}

func (t Test) GetName() string {
	return t.Name
}

func getstring(i interface{}) string {
	t, ok := i.(Test)
	if ok {
		return t.GetName()
	}
	return "Not this type!"
}

func reflectGetString(i interface{}) string {
	value := reflect.ValueOf(i)
	testI := value.Interface()
	result, ok := testI.(Test)
	if ok {
		return result.GetName()
	}
	return "Error!"

}

func main() {
	myTest := Test{Name: "Leo"}
	s := getstring(myTest)
	fmt.Println(s)
	s2 := reflectGetString(myTest)
	fmt.Println(s2)
}
