package main

import (
	"net/http"
	"tokener"
)

func main() {
	http.HandleFunc("/", GetRoot)
	http.HandleFunc("/reference", GetReference)
	http.HandleFunc("/token", GetToken)
	http.ListenAndServe(":8000", nil)
}
