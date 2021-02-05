package controller

import (
	"bytes"
	"encoding/json"
	"github.com/brunodecastro/digital-accounts/app/api/auth"
	"github.com/brunodecastro/digital-accounts/app/api/response"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/errors"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/service"
	"net/http"
	"testing"
)

func TestAuthenticationController_Authenticate_Success(t *testing.T) {
	t.Parallel()

	var credentialInputVO = input.CredentialInputVO{
		CPF:    "00801246156",
		Secret: "123456",
	}
	body, _ := json.Marshal(credentialInputVO)
	var credentialVO = vo.CredentialVO{
		CPF:       "00801246156",
		AccountID: "0001",
		Username:  "Bruno de Castro Oliveira",
	}
	controller := NewAuthenticationController(
		service.MockAuthenticationService{
			Result: credentialVO,
			Err:    nil,
		})

	wantStatusCode := http.StatusOK

	endPoint := "/login"
	req, _ := http.NewRequest(http.MethodPost, endPoint, bytes.NewReader(body))
	rec := mockRequestHandler(req, http.MethodPost, endPoint, true, controller.Authenticate)

	// Check the response status code
	if statusCode := rec.Code; statusCode != wantStatusCode {
		t.Errorf("Authenticate() error = %v, wantErr %v", wantStatusCode, statusCode)
	}
}

func TestAuthenticationController_Authenticate_Error_Without_Authorization(t *testing.T) {
	t.Parallel()

	var credentialInputVO = input.CredentialInputVO{
		CPF:    "00801246156",
		Secret: "123456",
	}
	body, _ := json.Marshal(credentialInputVO)

	wantStatusCode := http.StatusUnauthorized
	wantErr := true
	wantErrorMsg := custom_errors.ErrAuthorizationHeader.Error()

	endPoint := "/login"
	req, _ := http.NewRequest(http.MethodPost, endPoint, bytes.NewReader(body))
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {}
	rec := mockRequestHandler(req, http.MethodPost, endPoint, false, auth.AuthorizeMiddleware(handlerFunc))

	// Check the response status code
	if statusCode := rec.Code; statusCode != wantStatusCode {
		t.Errorf("Authenticate() error = %v, wantErr %v", wantStatusCode, statusCode)
	}

	// Check result response
	var responseError response.HTTPErrorResponse
	_ = json.Unmarshal(rec.Body.Bytes(), &responseError)

	if wantErr && responseError.Error.Message != wantErrorMsg {
		t.Errorf("Authenticate() error = %v, wantErr %v", responseError.Error.Message, wantErrorMsg)
		return
	}
}

func TestAuthenticationController_Authenticate_Error_Invalid_Token(t *testing.T) {
	t.Parallel()

	var credentialInputVO = input.CredentialInputVO{
		CPF:    "00801246156",
		Secret: "123456",
	}
	body, _ := json.Marshal(credentialInputVO)

	wantStatusCode := http.StatusUnauthorized
	wantErr := true
	wantErrorMsg := custom_errors.ErrInvalidToken.Error()

	endPoint := "/login"
	req, _ := http.NewRequest(http.MethodPost, endPoint, bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+"INVALID_TOKEN_SIMULATION")
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {}
	rec := mockRequestHandler(req, http.MethodPost, endPoint, false, auth.AuthorizeMiddleware(handlerFunc))

	// Check the response status code
	if statusCode := rec.Code; statusCode != wantStatusCode {
		t.Errorf("Authenticate() error = %v, wantErr %v", wantStatusCode, statusCode)
	}

	// Check result response
	var responseError response.HTTPErrorResponse
	_ = json.Unmarshal(rec.Body.Bytes(), &responseError)

	if wantErr && responseError.Error.Message != wantErrorMsg {
		t.Errorf("Authenticate() error = %v, wantErr %v", responseError.Error.Message, wantErrorMsg)
		return
	}
}
