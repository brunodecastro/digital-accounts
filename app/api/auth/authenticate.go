package auth

import (
	"github.com/brunodecastro/digital-accounts/app/common/logger"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/brunodecastro/digital-accounts/app/config"
	"go.uber.org/zap"
	"time"
)

// SignedTokenString return the token auth
func SignedTokenString(credentialOutput vo.CredentialVO) string {

	// Token expiration time
	var maxTokenLiveTime, err = time.ParseDuration(config.AppConfig.AuthConfig.MaxTokenLiveTime)
	if err != nil {
		logger.GetLogger().Error("error parsing the token max live time", zap.Error(err))
	}
	tokenExpirationTime := time.Now().Add(maxTokenLiveTime * time.Minute)

	credentialClaims := vo.CredentialClaimsVO{
		Username:  credentialOutput.Username,
		AccountID: credentialOutput.AccountID,
		ExpiresAt: tokenExpirationTime.Unix(),
	}

	return GenerateToken(credentialClaims)
}
