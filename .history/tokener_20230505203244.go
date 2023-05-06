package main

import (
	"net/http"
	"tokener/router"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/reference", getReference)
	http.HandleFunc("/token", getToken)
	http.ListenAndServe(":8000", nil)
}
