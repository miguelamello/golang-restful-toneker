package validation

import (
	"fmt"
	s "strings"
)

func ValidateBearerString(bearer string) bool {
	cBearer := s.Trim(bearer, " ")
	fmt.Println(len(cBearer))
	return true
}
