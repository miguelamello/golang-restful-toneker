package validation

import (
	"fmt"
	s "strings"
)

func ValidateBearer(bearer string) bool {
	cs.Trim(bearer, " ")
	return true
}
