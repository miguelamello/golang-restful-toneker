package validation

import (
	"fmt"
	s "strings"
)

func ValidateBearerString(bearer string) bool {
	cBearer := s.Trim(bearer, " ")
	if len(cBearer) == 0 { return false
	return true
}
