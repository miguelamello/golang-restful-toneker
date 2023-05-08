package validation

import (
	"encoding/json"
	"net/mail"
	s "strings"
)

type Response200 struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Payload string `json:"payload"`
}
type Response400 struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func is

// Verify if the email string is valid
func ValidateEmailString(email string) bool {
	cEmail := s.Trim(email, " ")
	if len(cEmail) == 0 { 
		return false 
	} else { 
		return true 
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

func FormatResponse200( status string, message string, payload string) string {
	response := Response200{
		Status: status,
		Message: message, 
		Payload: payload,
	}
	jsonBytes, _ := json.Marshal(response)
	jsonString := string(jsonBytes)
	return jsonString
}

func FormatResponse400( status string, message string) string {
	response := Response400{
		Status: status,
		Message: message, 
	}
	jsonBytes, _ := json.Marshal(response)
	jsonString := string(jsonBytes)
	return jsonString
}