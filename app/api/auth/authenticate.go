package auth

import (
	"github.com/brunodecastro/digital-accounts/app/common/constants"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

func SignedTokenString(credentialOutput vo.CredentialVO) string {

	// TODO: move maxTokenExpirationTime to config
	// Token expiration time (15 minutes)
	var maxTokenExpirationTime time.Duration = 15
	tokenExpirationTime := time.Now().Add(maxTokenExpirationTime * time.Minute)

	credentialClaims := vo.CredentialClaimsVO{
		Username:  credentialOutput.Username,
		AccountId: credentialOutput.AccountId,
		ExpiresAt: tokenExpirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, credentialClaims)
	tokenString, err := token.SignedString([]byte(constants.JwtSecretKey))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}
