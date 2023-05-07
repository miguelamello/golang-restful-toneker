package validation

import (
	"fmt"
	s "strings"
)

func ValidateBearer(bearer string) bool {
	s.Trim(bearer, " ")
	return true
}
