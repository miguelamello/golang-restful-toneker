package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", getDocumentation)
	http.HandleFunc("/token", getToken)
	http.ListenAndServe(":8000", nil)
}

func getDocumentation(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./documentationindex.html")
}

func getToken(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "{ token: 1234567890 }")
}
