package util

import (
	"github.com/brunodecastro/digital-accounts/app/util/constants"
	"regexp"
)

// NumbersOnly returns only numbers of string
func NumbersOnly(str string) string {
	reg, err := regexp.Compile("[^0-9]+")
	MaybeError(err, "error when compile only number regex")
	return reg.ReplaceAllString(str, "")
}

// FormatCpf format a cpf number
func FormatCpf(str string) string {
	reg, err := regexp.Compile(constants.CPFFormatPattern)
	MaybeError(err, "error when compile cpf format regex")
	return reg.ReplaceAllString(str, "$1.$2.$3-$4")
}
