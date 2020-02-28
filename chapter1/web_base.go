package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", nil)
}

// handler的要义就是，给一个Request，里面是请求信息，再给一个Writer，将某些信息给写进去。
func handler(w http.ResponseWriter, r *http.Request) {
	// Must write things into ResponseWriter
	fmt.Fprintf(w, "Request received %v", r.URL)
}
