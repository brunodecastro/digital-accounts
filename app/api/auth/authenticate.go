package auth

import (
	"github.com/brunodecastro/digital-accounts/app/common/constants"
	"github.com/brunodecastro/digital-accounts/app/common/logger"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"log"
	"time"
)

func SignedTokenString(credentialOutput vo.CredentialVO) string {

	// Token expiration time
	var maxTokenLiveTime, err = time.ParseDuration(config.AppConfig.AuthConfig.MaxTokenLiveTime)
	if err != nil {
		logger.GetLogger().Error("error parsing the token max live time", zap.Error(err))
	}
	tokenExpirationTime := time.Now().Add(maxTokenLiveTime * time.Minute)

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
