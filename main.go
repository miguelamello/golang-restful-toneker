package main

import (
	"net/http"
	"tokener/routes"
	"github.com/sirupsen/logrus"
)

func main() {

	// Routes
	http.HandleFunc("/", routes.GetRoot)
	http.HandleFunc("/reference", routes.GetReference)
	http.HandleFunc("/token", routes.GetToken)

	// Listen on port 3010
	http.ListenAndServe(":3010", nil)
	logrus.Info("Server started on port 3010")
	
}
