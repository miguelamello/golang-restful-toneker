package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/", getToken)
	http.ListenAndServe(":8000", nil)
}

func getRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Documentation")
}

func getToken(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "[]")
}
