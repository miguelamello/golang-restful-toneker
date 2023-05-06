package main

import (
	"net/http"
	"tokener/routes"
)

func main() {
	
	http.HandleFunc("/", routes.GetRoot)
	http.HandleFunc("/reference", routes.GetReference)
	http.HandleFunc("/token", routes.GetToken)

	// Catchall handler for non-existent routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	http.ListenAndServe(":8000", nil)
}


