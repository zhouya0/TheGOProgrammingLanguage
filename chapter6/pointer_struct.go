package main

import (
	"fmt"
)

type Animal interface {
	Talk()
	Walk()
}

type Cat struct {
	Age int
}

func (c Cat) Walk() {
	c.Age++
	fmt.Println("cat walking")
	fmt.Println(c.Age)
}

func (c Cat) Talk() {
	fmt.Println("cat talking")
	fmt.Println(c.Age)
}

func main() {
	var c Animal = &Cat{Age: 3}
	c.Walk() // copied
	c.Talk() // copied
}
