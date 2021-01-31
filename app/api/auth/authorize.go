package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/brunodecastro/digital-accounts/app/api/response"
	"github.com/brunodecastro/digital-accounts/app/common/constants"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func AuthorizeMiddleware(next http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := parseBearerToken(bearerToken[1])
				if err != nil {
					response.CreateErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}
				if token.Valid {
					claims, _ := token.Claims.(jwt.MapClaims)
					credentialClaims := vo.CredentialClaimsVO{
						AccountId: claims["account_id"].(string),
						Username:  claims["username"].(string),
					}
					// create a new context with credential claims value
					newContext := context.WithValue(req.Context(), constants.CredentialClaimsContextKey, credentialClaims)
					next(w, req.WithContext(newContext))
					return
				}
				json.NewEncoder(w).Encode(Exception{Message: "Invalid Authorization token"})
				return
			}
		}
		json.NewEncoder(w).Encode(Exception{Message: "An Authorization header is required"})
	}
}

func parseBearerToken(bearerToken string) (*jwt.Token, error) {
	return jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(constants.JwtSecretKey), nil
	})
}

type Exception struct {
	Message string `json:"message"`
}
