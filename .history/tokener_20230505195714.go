package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8001", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Tokener okay")
}