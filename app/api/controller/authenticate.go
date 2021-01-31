package controller

import (
	"encoding/json"
	"github.com/brunodecastro/digital-accounts/app/api/response"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

var secretKey string

func init() {
	secretKey = "JWT_KEY" //TODO: pegar do config
}

type AuthenticateController struct {
	service service.AuthenticateService
}

func NewAuthenticateController(service service.AuthenticateService) AuthenticateController {
	return AuthenticateController{
		service: service,
	}
}

func (controller AuthenticateController) Authenticate(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var credentialInput input.CredentialInputVO
	err := json.NewDecoder(req.Body).Decode(&credentialInput)
	if err != nil {
		response.CreateErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// Validate credentials
	credentialOutputVO, err := controller.service.Authenticate(req.Context(), credentialInput)
	if err != nil {
		response.CreateErrorResponse(w, http.StatusForbidden, "invalid credentials")
		return
	}

	response.CreateSuccessResponse(w, http.StatusOK, output.CredentialOutputVO{
		Token: signedTokenString(credentialOutputVO),
	})
	return
}

func signedTokenString(credentialOutput output.CredentialOutputVO) string {

	// Token expiration time (15 minutes)
	var maxTokenExpirationTime time.Duration = 15
	tokenExpirationTime := time.Now().Add(maxTokenExpirationTime * time.Minute)

	credentialClaims := vo.CredentialClaims{
		Username:  credentialOutput.Username,
		AccountId: credentialOutput.AccountId,
		ExpiresAt: tokenExpirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, credentialClaims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}
