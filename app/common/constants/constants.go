package constants

import "time"

// common api constants
const (
	ProfileDev                 string = "dev"
	ProfileProd                string = "prod"
	CPFFormatPattern           string = `([\d]{3})([\d]{3})([\d]{3})([\d]{2})`
	DateDefaultLayout          string = time.RFC3339 // "2006-01-02T15:04:05Z07:00"
	JSONContentType            string = "application/json"
	HTMLContentType            string = "text/html; charset=utf-8"
	CredentialClaimsContextKey string = "credentialClaims"
	TransactionContextKey      int    = iota
)
