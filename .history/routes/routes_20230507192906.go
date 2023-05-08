package routes

import (
  "fmt"
  "net/http"
  "tokener/database"
  "tokener/validation"
)

// Route handler for showing the Service Presentation
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

// Route handler for showing the Reference Documentation
func GetReference(res http.ResponseWriter, req *http.Request) {
  http.ServeFile(res, req, "./html/reference/index.html")
}

// Route handler for getting the registered token or creating a new one
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

  // Validate the email string from the header variables
  user := validation.ValidateEmailString(req.Header.Get("user"))
  if !user {
    response := validation.FormatResponse400(
      "400", 
      "User not valid. Verify user email presence in header variables.", 
    )
    fmt.Fprint(res, response)
    return
  }

  payload := database.GetPayload(req.Header.Get("user"))

  if len(payload) <= 2 {
      // Return the token related to the user mail
      database.CreateItem(req.Header.Get("user"))
      payload := database.GetPayload(req.Header.Get("user"))
      if len(payload) <= 2 {
          response := validation.FormatResponse400(
              "400", 
              "A token was not created. Try again later.", 
          )
          fmt.Fprint(res, response)
          return
      }
  }
  fmt.Fprint(res, payload)
  
  
  // Return the token related to the user mail
  //database.GetItem("webmaster@orionsoft.site")
  //json.NewEncoder(res).Encode(token)

  /*response := validation.FormatResponse200(
      "200", 
      "Token found", 
      `payload: { 
          user: 'miguelangelomello@gmail.com', 
          token: 1234567890,
          created: '2023-01-01 00:00:00',
          expires: '2018-02-01 00:00:00'
      }`,
  )
  fmt.Fprint(res, response)*/
}

