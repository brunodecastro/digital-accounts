package constants

import "time"

const (
	ProfileDev                 string = "dev"
	ProfileProd                string = "prod"
	CPFFormatPattern           string = `([\d]{3})([\d]{3})([\d]{3})([\d]{2})`
	DateDefaultLayout          string = time.RFC3339 // "2006-01-02T15:04:05Z07:00"
	JsonContentType            string = "application/json"
	JWTSecretKey               string = "jwt-secret-key" //TODO: move to config
	CredentialClaimsContextKey string = "credentialClaims"
	TransactionContextKey      int    = iota
)
