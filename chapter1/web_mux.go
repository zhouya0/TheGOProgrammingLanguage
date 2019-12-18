package main

import (
	"fmt"
	"net/http"
	"sync"
)

var mu sync.Mutex
var n int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", count)
	http.ListenAndServe("localhost:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	n++
	fmt.Fprintf(w, "request %v recieved", r.URL)
	mu.Unlock()
}

func count(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "already received %d requests", n)
	mu.Unlock()
}
