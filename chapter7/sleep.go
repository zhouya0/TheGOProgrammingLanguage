package main

import (
	"flag"
	"fmt"
	"time"
)

// 参数，默认值，注解
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}