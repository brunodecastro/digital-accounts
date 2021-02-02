package controller

import (
	"encoding/json"
	"github.com/brunodecastro/digital-accounts/app/api/auth"
	"github.com/brunodecastro/digital-accounts/app/api/response"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/errors"
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

// Create godoc
// @Summary Creates a new transfer
// @Description Creates a new transfer
// @tags Transfers
// @Accept  json
// @Produce  json
// @Param account body input.CreateTransferInputVO true "Transfer Input"
// @Success 201 {object} output.CreateTransferOutputVO
// @Failure 400,500 {object} response.HTTPErrorResponse
// @Security ApiKeyAuth
// @Router /transfers [post]
//
// Create - creates a new transfer
func (controller TransferController) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var transferInputVO input.CreateTransferInputVO
	if err := json.NewDecoder(r.Body).Decode(&transferInputVO); err != nil {
		response.CreateErrorResponse(w, http.StatusBadRequest, custom_errors.ErrorInvalidJSONFormat.Error())
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

// FindAll godoc
// @Summary List all transfers
// @Description List all transfers
// @tags Transfers
// @Produce  json
// @Success 200 {object} output.FindAllTransferOutputVO
// @Failure 500 {object} response.HTTPErrorResponse
// @Security ApiKeyAuth
// @Router /transfers [get]
//
// FindAll list all transfers
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
