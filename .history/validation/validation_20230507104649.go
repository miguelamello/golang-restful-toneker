package validation

import (
	"encoding/json"
	s "strings"
)

type Response400 struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

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

func FormatResponse400( status string, message string) string {
	response := Response400{
		Status: "400",
		Message: "Email not valid. Verify email presence in header variables.", 
	}
	jsonBytes, _ := json.Marshal(response)
	jsonString := string(jsonBytes)
}
