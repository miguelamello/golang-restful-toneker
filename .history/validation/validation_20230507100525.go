package validation

import (
	s "strings"
)

// Verify if the bearer string is valid
func validateEmailString(email string) bool {
	cBearer := s.Trim(email, " ")
	if len(cBearer) == 0 { 
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