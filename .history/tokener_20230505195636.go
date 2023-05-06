package main

import (
	"fmt"
	"net/http"
)

func main() {
    
}

func hello(w http.ResponseWriter, req *http.Request) {
	http.HandleFunc("/hello", hello)
}