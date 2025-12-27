package validator

import "regexp"

var lowerLatinRegex = regexp.MustCompile(`^[a-z]+$`)

func IsValidKey(key string) bool {
	return lowerLatinRegex.MatchString(key)
}
