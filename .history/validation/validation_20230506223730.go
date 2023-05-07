package validation

import (
	s "strings"
)

func ValidateBearerString(bearer string) bool {
	cBearer := s.Trim(bearer, " ")
	if len(cBearer) == 0 { 
		return false 
	} else { 
		return true 
	}
}
