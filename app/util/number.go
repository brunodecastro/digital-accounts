package util

import (
	"regexp"
)

// NumbersOnly returns only numbers of string
func NumbersOnly(str string) string {
	reg, err := regexp.Compile("[^0-9]+")
	MaybeError(err, "error when compile only number regex")
	return reg.ReplaceAllString(str, "")
}
