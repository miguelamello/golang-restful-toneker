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
        // Dont return a error but tell client that the route is not found
        fmt.Fprintf(res, "{ status: 404, message: 'Route Not Found' }")
    }
}

func GetReference(res http.ResponseWriter, req *http.Request) {
    http.ServeFile(res, req, "./html/reference/index.html")
}

func GetToken(res http.ResponseWriter, req *http.Request) {
    // Set the Content-Type header to application/json
    res.Header().Set("Content-Type", "application/json")
    // Return the token related to the user mail
    fmt.Fprintf(res, "{ status: 200token: 1234567890 }")
}
