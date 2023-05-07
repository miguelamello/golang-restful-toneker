package validation

import (
	"fmt"
	s "strings"
)

func ValidateBearerString(bearer string) bool {
	cBearer := s.Trim(bearer, " ")
	len(cBearer)
	return true
}
