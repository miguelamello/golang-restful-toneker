package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/documentation", getReference)
	http.HandleFunc("/token", getToken)
	http.ListenAndServe(":8000", nil)
}

func get(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./documentation/index.html")
}

func getDocumentation(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./documentation/index.html")
}

func getToken(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "{ token: 1234567890 }")
}
