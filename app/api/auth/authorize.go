package auth

import (
	"context"
	"fmt"
	"github.com/brunodecastro/digital-accounts/app/api/response"
	"github.com/brunodecastro/digital-accounts/app/common/constants"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/errors"
	"github.com/brunodecastro/digital-accounts/app/common/logger"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// AuthorizeMiddleware middleware to authorize the access to resource
func AuthorizeMiddleware(next http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := parseBearerToken(bearerToken[1])
				if err != nil {
					logger.GetLogger().Error("error parsing the token", zap.Error(err))
					response.CreateErrorResponse(w, http.StatusUnauthorized, custom_errors.ErrInvalidToken.Error())
					return
				}
				if token.Valid {
					claims, _ := token.Claims.(jwt.MapClaims)
					credentialClaims := vo.CredentialClaimsVO{
						AccountID: claims["account_id"].(string),
						Username:  claims["username"].(string),
					}
					// create a new context with credential claims value
					newContext := context.WithValue(req.Context(), constants.CredentialClaimsContextKey, credentialClaims)
					next(w, req.WithContext(newContext))
					return
				}
				response.CreateErrorResponse(w, http.StatusUnauthorized, custom_errors.ErrInvalidAuthorizationToken.Error())
				return
			}
		}
		response.CreateErrorResponse(w, http.StatusUnauthorized, custom_errors.ErrAuthorizationHeader.Error())
	}
}

// parse the bearer jwt token
func parseBearerToken(bearerToken string) (*jwt.Token, error) {
	return jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error on jwt parse")
		}
		return []byte(config.GetAPIConfigs().AuthConfig.JWTSecretKey), nil
	})
}
