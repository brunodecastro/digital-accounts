package util

import (
	"github.com/brunodecastro/digital-accounts/app/common/constants"
	"time"
)

// FormatDate format date with default layout
func FormatDate(date time.Time) string {
	return FormatDateLayout(date, constants.DateDefaultLayout)
}

// FormatDatePattern format date with layout
func FormatDateLayout(date time.Time, layout string) string {
	return date.Format(layout)
}
