package controller

import (
	"encoding/json"
	"github.com/brunodecastro/digital-accounts/app/api/auth"
	"github.com/brunodecastro/digital-accounts/app/api/response"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/custom-errors"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AuthenticationController struct {
	service service.AuthenticationService
}

func NewAuthenticationController(service service.AuthenticationService) AuthenticationController {
	return AuthenticationController{
		service: service,
	}
}

func (controller AuthenticationController) Authenticate(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var credentialInput input.CredentialInputVO
	err := json.NewDecoder(req.Body).Decode(&credentialInput)
	if err != nil {
		response.CreateErrorResponse(w, http.StatusBadRequest, custom_errors.ErrorInvalidJsonFormat.Error())
		return
	}

	// Validate credentials
	credentialOutputVO, err := controller.service.Authenticate(req.Context(), credentialInput)
	if err != nil {
		response.CreateErrorResponse(w, http.StatusForbidden, custom_errors.ErrInvalidAccessCredentials.Error())
		return
	}

	response.CreateSuccessResponse(w, http.StatusOK, output.CredentialOutputVO{
		Token: auth.SignedTokenString(credentialOutputVO),
	})
	return
}
