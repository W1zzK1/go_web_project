package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go!")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe("localhost:8080", nil)
}
func main() {
	handleRequest()
}
