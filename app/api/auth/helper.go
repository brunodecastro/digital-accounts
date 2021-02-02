package auth

import (
	"github.com/brunodecastro/digital-accounts/app/common/constants"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

// GetAccountIdFromAuth return the account id from the auth token
func GetAccountIdFromToken(req *http.Request) string {
	// Get the credential claims from the context
	credentialClaimsVO := req.Context().Value(constants.CredentialClaimsContextKey).(vo.CredentialClaimsVO)
	return credentialClaimsVO.AccountId
}

func GenerateToken(credentialClaims vo.CredentialClaimsVO) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, credentialClaims)
	tokenString, err := token.SignedString([]byte(constants.JwtSecretKey))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}

func SetContextValue(req *http.Request) {
	/**
	http.HandlerFunc()

	// create a new context with credential claims value
	newContext := context.WithValue(req.Context(), constants.CredentialClaimsContextKey, credentialClaims)
	next(w, req.WithContext(newContext))
	*/

}
