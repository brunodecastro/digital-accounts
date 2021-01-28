package constants

import "time"

const (
	ProfileDev               = "dev"
	ProfileProd              = "prod"
	CPFFormatPattern  string = `([\d]{3})([\d]{3})([\d]{3})([\d]{2})`
	DateDefaultLayout string = time.RFC3339 // "2006-01-02T15:04:05Z07:00"
	JSON_CONTENT_TYPE        = "application/json"
)
