package routes

import (
	"fmt"
	"net/http"
)

func GetRoot(res http.ResponseWriter, req *http.Request) {
    if req.URL.Path == "/" {
        http.ServeFile(res, req, "./html/root/index.html")
    } else {
        // Set the Content-Type header to application/json
		res.Header().Set("Content-Type", "application/json")
        
        fmt.Fprintf(res, "{ status: 404, message: 'Route Not Found' }")
    }
}

func GetReference(res http.ResponseWriter, req *http.Request) {
    http.ServeFile(res, req, "./html/reference/index.html")
}

func GetToken(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "{ token: 1234567890 }")
}
