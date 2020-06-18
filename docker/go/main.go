package main

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func main() {
	fmt.Printf("Starting server at 'localhost:8082'\n")

	http.HandleFunc("/", test)
	http.ListenAndServe(":8082", nil)
}
