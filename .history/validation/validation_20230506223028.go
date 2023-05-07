package validation

import (
	"fmt"
	s "strings"
)

func ValidateBearer(bearer string) bool {
	cBearer := s.Trim(bearer, " ")
	fmt.Println(len(cBearer)
	return true
}
