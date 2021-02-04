package util

import (
	"fmt"
	"github.com/brunodecastro/digital-accounts/app/common/constants"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

var (
	cpfFirstDigitTable  = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	cpfSecondDigitTable = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
)

// IsCpfValid - returns if CPF is a valid CPF document
func IsCpfValid(cpfStr string) bool {
	return validateCPF(NumbersOnly(cpfStr))
}

// ValidateCPF validates a CPF document
func validateCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	firstPart := cpf[0:9]
	sum := sumDigit(firstPart, cpfFirstDigitTable)

	r1 := sum % 11
	d1 := 0

	if r1 >= 2 {
		d1 = 11 - r1
	}

	secondPart := firstPart + strconv.Itoa(d1)

	dsum := sumDigit(secondPart, cpfSecondDigitTable)

	r2 := dsum % 11
	d2 := 0

	if r2 >= 2 {
		d2 = 11 - r2
	}

	finalPart := fmt.Sprintf("%s%d%d", firstPart, d1, d2)
	return finalPart == cpf
}

func sumDigit(s string, table []int) int {
	if len(s) != len(table) {
		return 0
	}

	sum := 0

	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}
	return sum
}

// FormatCpf format a cpf number
func FormatCpf(cpfString string) string {
	reg, err := regexp.Compile(constants.CPFFormatPattern)
	MaybeError(err, "error when compile cpf format regex")
	return reg.ReplaceAllString(cpfString, "$1.$2.$3-$4")
}

// GenerateCPFOnlyNumbers - generate a cpf with only numbers
func GenerateCPFOnlyNumbers() (cpfString string) {
	rand.Seed(time.Now().UTC().UnixNano())
	cpf := rand.Perm(9)
	cpf = append(cpf, verify(cpf, len(cpf)))
	cpf = append(cpf, verify(cpf, len(cpf)))

	for _, c := range cpf {
		cpfString += strconv.Itoa(c)
	}
	return
}

// GenerateCPFFormatted - generate a cpf formatted
func GenerateCPFFormatted() (cpf string) {
	return FormatCpf(GenerateCPFOnlyNumbers())
}

func verify(data []int, n int) int {
	var total int

	for i := 0; i < n; i++ {
		total += data[i] * (n + 1 - i)
	}

	total = total % 11
	if total < 2 {
		return 0
	}
	return 11 - total
}
