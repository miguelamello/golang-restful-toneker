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

	// Listen on port 8000
	http.ListenAndServe(":8000", nil)
	
}


