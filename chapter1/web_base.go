package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Must write things into ResponseWriter!!!
	// Must write things into ResponseWriter!!!
	// Must write things into ResponseWriter!!!
	fmt.Fprintf(w, "Request received %v", r.URL)
}
