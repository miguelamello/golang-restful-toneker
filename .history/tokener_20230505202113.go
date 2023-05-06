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

func getRoot(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./routes/root/index.html")
}

func getReference(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./routes/reference/index.html")
}

func getToken(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "{ token: 1234567890 }")
}
