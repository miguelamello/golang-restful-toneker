package validation

import (
	"fmt"
	s "strings"
)

func ValidateBearer(bearer string) bool {
	fmt.Println(s.Trim(bearer))
	return true
}
