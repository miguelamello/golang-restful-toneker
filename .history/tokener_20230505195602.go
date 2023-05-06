package main

import (
	"fmt"
	"net/http"
)

func main() {
    fmt.Println("Tokener okay")
}

func hello(w http.ResponseWriter, req *http.Request) {
	
