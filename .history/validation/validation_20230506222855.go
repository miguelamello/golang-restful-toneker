package validation

import (
	"fmt"
	s "strings"
)

func ValidateBearer(bearer string) bool {
	cBeas.Trim(bearer, " ")
	return true
}
