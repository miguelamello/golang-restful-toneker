package main

import (
	"net/http"
	"tokener/routes"
)

func main() {
	// Routes
	http.HandleFunc("/", routes.GetRoot)
	http.HandleFunc("/reference", routes.GetReference)
	http.HandleFunc("/token", routes.GetToken)

	// Listen 
	http.ListenAndServe(":8000", nil)
}


