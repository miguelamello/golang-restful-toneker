package validation

import (
	"encoding/json"
	"fmt"
	//"fmt"
	"regexp"
	s "strings"
)

// Struct to hold payload data
type Payload struct {
	Email   string `json:"email"`
	Token   string `json:"token"`
	Created string `json:"created"`
	Expires string `json:"expires"`
}

// Struct to hold 200 response data
type Response200 struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Payload Payload `json:"payload"`
}

// Struct to hold 400 response data
type Response400 struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Verify if the email string is valid
func isEmail(email string) bool {
	// Define a regular expression for validating email addresses
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	// Check if the email address matches the regular expression
	if re.MatchString(email) {
		return true
	} else {
		return false
	}
}

// Verify if the email string exists and is valid
func ValidateEmailString(email string) bool {
	cEmail := s.Trim(email, " ")
	if len(cEmail) == 0 { 
		return false 
	} else { 
		if isEmail(cEmail) {
			return true 
		} else {			
			return false
		}
	}
}

// Verify if the bearer string is valid
func ValidateBearerString(bearer string) bool {
	cBearer := s.Trim(bearer, " ")
	if len(cBearer) == 0 { 
		return false 
	} else { 
		return true 
	}
}

// Format the response for 200 responses
func FormatResponse200( status string, message string, payload string) string {
	p := Payload{} // Create an empty payload struct
	fmt.Printf(p)
	json.Unmarshal([]byte(payload), &p) // Unmarshal the payload string into the payload struct
	response := Response200{
		Status: status,
		Message: message, 
		Payload: p,
	} // Create a response struct with the payload struct
	jsonBytes, _ := json.Marshal(response) // Marshal the response struct into a json string
	jsonString := string(jsonBytes) // Convert the json bytes into a string
	return jsonString
}

// Format the response for 400 responses
func FormatResponse400( status string, message string) string {
	response := Response400{
		Status: status,
		Message: message, 
	}
	jsonBytes, _ := json.Marshal(response)
	jsonString := string(jsonBytes)
	return jsonString
}
