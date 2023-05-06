package main

import (
	"net/http"
	"tokener/routes"
)

func main() {
	http.HandleFunc("/", routes.GetRoot)
	http.HandleFunc("/reference", routes.GetReference)
	http.HandleFunc("/token", routes.GetToken)

	
	http.ListenAndServe(":8000", nil)
}


