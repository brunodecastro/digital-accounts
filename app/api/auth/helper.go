package auth

import (
	"github.com/brunodecastro/digital-accounts/app/common/constants"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

// GetAccountIDFromToken - return the account id from the auth token
func GetAccountIDFromToken(req *http.Request) string {
	// Get the credential claims from the context
	credentialClaimsVO := req.Context().Value(constants.CredentialClaimsContextKey).(vo.CredentialClaimsVO)
	return credentialClaimsVO.AccountID
}

// GenerateToken generates a jwt token with claims information
func GenerateToken(credentialClaims vo.CredentialClaimsVO) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, credentialClaims)
	tokenString, err := token.SignedString([]byte(constants.JWTSecretKey))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}
