package util

import (
	"regexp"
)

func NumbersOnly(s string) (string) {
	reg, err := regexp.Compile("[^0-9]+")
	MaybeError(err, "error when compile only number regex")
	return reg.ReplaceAllString(s, "")
}
