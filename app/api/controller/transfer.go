package controller

import (
	"encoding/json"
	"github.com/brunodecastro/digital-accounts/app/api/auth"
	"github.com/brunodecastro/digital-accounts/app/api/response"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/custom-errors"
	"github.com/brunodecastro/digital-accounts/app/common/validator"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/service"
	"net/http"
)

// TransferController - struct of Transfer Controller
type TransferController struct {
	service service.TransferService
}

// NewTransferController - new Transfer Controller instance
func NewTransferController(service service.TransferService) TransferController {
	return TransferController{
		service: service,
	}
}

// Create - creates a new account
func (controller TransferController) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var transferInputVO input.CreateTransferInputVO
	if err := json.NewDecoder(r.Body).Decode(&transferInputVO); err != nil {
		response.CreateErrorResponse(w, http.StatusBadRequest, custom_errors.ErrorInvalidJsonFormat.Error())
		return
	}

	// Validate input fields
	if err := validator.ValidateCreateTransferInput(transferInputVO); err != nil {
		response.CreateErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// get the account origin id from the auth token
	transferInputVO.AccountOriginID = auth.GetAccountIDFromToken(r)

	transferCreated, err := controller.service.Create(r.Context(), transferInputVO)
	if err != nil {
		response.CreateErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.CreateSuccessResponse(w, http.StatusCreated, transferCreated)
}

// FindAll - list all accounts
func (controller TransferController) FindAll(w http.ResponseWriter, r *http.Request) {

	// get the account id from the auth token
	accountOriginID := auth.GetAccountIDFromToken(r)

	transfers, err := controller.service.FindAll(r.Context(), accountOriginID)
	if err != nil {
		response.CreateErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.CreateSuccessResponse(w, http.StatusOK, transfers)
}
