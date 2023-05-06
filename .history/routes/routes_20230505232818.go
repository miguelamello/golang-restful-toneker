package routes

import (
	"fmt"
	"net/http"
)

func GetRoot(res http.ResponseWriter, req *http.Request) {
    http.ServeFile(res, req, "./html/root/index.html")
    http.NotFound(w, r)
}

func GetReference(res http.ResponseWriter, req *http.Request) {
    http.ServeFile(res, req, "./html/reference/index.html")
}

func GetToken(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "{ token: 1234567890 }")
}
