package vo

import "github.com/dgrijalva/jwt-go"

type CredentialClaimsVO struct {
	Username  string `json:"username"`
	AccountId string `json:"account_id"`
	ExpiresAt int64  `json:"expires_at"`
	jwt.StandardClaims
}
