package main

import "fmt"

var test = "fdsafdsgregaefcgsfgsfgs"

func main() {
	stringList := []byte(test)
	char := stringList[0]
	for _,v := range stringList {
		if v == char {
			continue
		} else {
			fmt.Println("Not Right!")
			return
		}
	}
	fmt.Println("Right")

}