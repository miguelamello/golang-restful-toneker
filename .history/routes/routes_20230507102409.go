package routes

import (
	"fmt"
	"net/http"
    "encoding/json"
    //"tokener/database"
    "tokener/validation"
)

type Response400 struct {
    status  string `json:"status"`
    message string `json:"message"`
}

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

    // Validate the bearer string from the header variables 
    // and return a error if it is not valid
    /*valid := validation.ValidateBearerString(req.Header.Get("bearerToken"))
    if !valid {
        fmt.Fprintf(res, `{ 
            status: 400, 
            message: 'bearerToken not valid. Verify bearerToken presence 
            in header variables or request a new one. 
            See Reference documentation.' 
        }`)
        return
    }*/

    email := validation.ValidateEmailString(req.Header.Get("email"))
    if !email {
        response := Response400{
            status: "400",
            message: 
        }
        fmt.Fprintf(res, `{ 
            status: 400, 
            message: 'Email not valid. Verify email presence 
            in header variables. 
            See Reference documentation.' 
        }`)
        return
    }
    
    // Return the token related to the user mail
    //database.CreateToken("miguelangelomello@gmail.com")
    // Return the token related to the user mail
    //database.GetItem("webmaster@orionsoft.site")
    //json.NewEncoder(res).Encode(token)
    fmt.Fprintf(res, `{ 
        status: 200, 
        message: 'Token found', 
        payload: { 
            user: 'miguelangelomello@gmail.com', 
            token: 1234567890,
            created: '2023-01-01 00:00:00',
            expires: '2018-02-01 00:00:00'
        }
    }`)
}