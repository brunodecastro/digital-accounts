package vo

import "github.com/dgrijalva/jwt-go"

// CredentialClaimsVO - vo that represents the credentials claims contained in the token auth
type CredentialClaimsVO struct {
	Username  string `json:"username"`
	AccountID string `json:"account_id"`
	jwt.StandardClaims
}
