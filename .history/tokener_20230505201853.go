package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", getDocumentation)
	http.HandleFunc("/documentation", getDocumentation)
	http.HandleFunc("/token", getToken)
	http.ListenAndServe(":8000", nil)
}

func getDocumentation(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./documentation/index.html")
}

func getToken(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "{ token: 1234567890 }")
}
