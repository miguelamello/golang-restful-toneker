package main

import (
	"net/http"
	"tokener/router"
)

func main() {
	http.HandleFunc("/", GetRoot)
	http.HandleFunc("/reference", getReference)
	http.HandleFunc("/token", getToken)
	http.ListenAndServe(":8000", nil)
}
