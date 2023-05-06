package main

import (
	"net/http"
	"tokener/routes"
)

func main() {
	http.HandleFunc("/", routes.GetRoot)
	http.HandleFunc("/reference", GetReference)
	http.HandleFunc("/token", GetToken)
	http.ListenAndServe(":8000", nil)
}
