package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	bs := make([]byte, 100)
	file.Seek(8,0)
	count, _ := file.Read(bs)
	fmt.Println(string(bs))
	fmt.Println(string(bs[:count]))
	fmt.Printf("count is: %v \n", count)

	// validate buff read line
	var eol = []byte{'\n'}
	fmt.Printf("The eol legth must be 1 right? %v \n", len(eol))
	file, _ = os.Open("./test.txt")
	r := bufio.NewReader(file)
	for i:=1; i<4;i++ {
		l, _ := r.ReadBytes(eol[0])
		fmt.Printf("This %v line: %v\n", i, string(l))
	}
}

//➜ chapter7 git:(master) ✗ go run seek.go
//aatest!
//Helloaworld!
//aatest!
//Helloaworld!
//count is: 20
//可以看到，这里的起点就是第八个字符了。

