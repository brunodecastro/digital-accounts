package controller

import (
	"encoding/json"
	"github.com/brunodecastro/digital-accounts/app/api/response"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/errors"
	"github.com/brunodecastro/digital-accounts/app/common/validator"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// AccountController - struct of Account Controller
type AccountController struct {
	service service.AccountService
}

// NewAccountController - new Account Controller instance
func NewAccountController(service service.AccountService) AccountController {
	return AccountController{
		service: service,
	}
}

// Create godoc
// @Summary Cria uma conta
// @Description Cria uma conta
// @tags Accounts
// @Accept  json
// @Produce  json
// @Param account body input.CreateAccountInputVO true "Dados da Conta"
// @Success 201 {object} output.CreateAccountOutputVO
// @Failure 400 {object} response.HTTPErrorResponse
// @Failure 500 {object} response.HTTPErrorResponse
// @Router /accounts [post]
//
// Create - creates a new account
func (controller AccountController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	var accountInputVO input.CreateAccountInputVO
	if err := json.NewDecoder(r.Body).Decode(&accountInputVO); err != nil {
		response.CreateErrorResponse(w, http.StatusBadRequest, custom_errors.ErrorInvalidJSONFormat.Error())
		return
	}

	// Validate input fields
	if err := validator.ValidateCreateAccountInput(accountInputVO); err != nil {
		response.CreateErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	accountCreated, err := controller.service.Create(r.Context(), accountInputVO)
	if err != nil {
		response.CreateErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.CreateSuccessResponse(w, http.StatusCreated, accountCreated)
}

// FindAll godoc
// @Summary Obtém a lista de contas
// @Description Obtém a lista de contas
// @tags Accounts
// @Produce  json
// @Success 200 {object} output.FindAllAccountOutputVO
// @Failure 500 {object} response.HTTPErrorResponse
// @Router /accounts [get]
//
// FindAll list all accounts
func (controller AccountController) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	accounts, err := controller.service.FindAll(r.Context())
	if err != nil {
		response.CreateErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.CreateSuccessResponse(w, http.StatusOK, accounts)
}

// GetBalance godoc
// @Summary Obtém o saldo da conta
// @Description Obtém o saldo da conta
// @tags Accounts
// @Accept  json
// @Produce  json
// @Param account_id path string true "ID da conta"
// @Success 200 {object} output.FindAccountBalanceOutputVO
// @Failure 400 {object} response.HTTPErrorResponse
// @Failure 500 {object} response.HTTPErrorResponse
// @Router /account/{account_id}/balance [get]
//
// GetBalance - Gets the account balance
func (controller AccountController) GetBalance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var accountID = params.ByName("account_id")

	// Validate input fields
	if err := validator.ValidateFindAccountBalanceInput(accountID); err != nil {
		response.CreateErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	account, err := controller.service.GetBalance(r.Context(), accountID)
	if err != nil {
		response.CreateErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.CreateSuccessResponse(w, http.StatusOK, account)
}
