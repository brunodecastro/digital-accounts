package controller

import (
	"encoding/json"
	"github.com/brunodecastro/digital-accounts/app/api/auth"
	"github.com/brunodecastro/digital-accounts/app/api/response"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/errors"
	"github.com/brunodecastro/digital-accounts/app/common/validator"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// AuthenticationController - struct of Authentication Controller
type AuthenticationController struct {
	service service.AuthenticationService
}

// NewAuthenticationController - new AuthenticationController instance
func NewAuthenticationController(service service.AuthenticationService) AuthenticationController {
	return AuthenticationController{
		service: service,
	}
}

// Authenticate godoc
// @Summary authenticate the user
// @Description authenticate the user in the api
// @tags Authentication
// @Accept  json
// @Produce  json
// @Param credential body input.CredentialInputVO true "Credential Input"
// @Success 201 {object} output.CreateTransferOutputVO
// @Failure 400 {object} response.HTTPErrorResponse
// @Failure 403 {object} response.HTTPErrorResponse
// @Security ApiKeyAuth
// @Router /login [post]
//
// Authenticate - authenticate the user in the api
func (controller AuthenticationController) Authenticate(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var credentialInput input.CredentialInputVO
	err := json.NewDecoder(req.Body).Decode(&credentialInput)
	if err != nil {
		response.CreateErrorResponse(w, http.StatusBadRequest, custom_errors.ErrorInvalidJSONFormat.Error())
		return
	}

	// Validate input fields
	if err := validator.ValidateAuthenticate(credentialInput); err != nil {
		response.CreateErrorResponse(w, http.StatusBadRequest, err.Error())
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
