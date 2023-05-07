package routes

import (
	"fmt"
	"net/http"
    //"tokener/database"
    "tokener/validation"
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

    valid := validation.ValidateBearerString(req.Header.Get("bearerToken"))
    if !valid {
        fmt.Fprintf(res, "{ status: 400, message: 'bearerToken not valid. Verify bearerToken presence in header variables or ' }")
        return
    }
    // Return the token related to the user mail
    //database.CreateToken("miguelangelomello@gmail.com")
    // Return the token related to the user mail
    //database.GetItem("webmaster@orionsoft.site")
    //json.NewEncoder(res).Encode(token)
    fmt.Fprintf(res, "{ status: 200, message: 'Token found', payload: { user: 'miguelangelomello@gmail.com', token: 1234567890} }")
}
