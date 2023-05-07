package validation

import (
	"fmt"
	s "strings"
)

func ValidateBearer(bearer string) bool {
	cBearer := s.Trim(bearer, " ")
	return true
}
