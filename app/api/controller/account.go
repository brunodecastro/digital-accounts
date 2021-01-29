package controller

import (
	"encoding/json"
	"github.com/brunodecastro/digital-accounts/app/api/response"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AccountController struct {
	service service.AccountService
}

func NewAccountController(service service.AccountService) AccountController {
	return AccountController{
		service: service,
	}
}

func (controller AccountController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	var accountInputVO input.CreateAccountInputVO
	if err := json.NewDecoder(r.Body).Decode(&accountInputVO); err != nil {
		response.CreateErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// TODO: validate fields

	accountCreated, err := controller.service.Create(r.Context(), accountInputVO)
	if err != nil {
		response.CreateErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.CreateSuccessResponse(w, http.StatusCreated, accountCreated)
}

func (controller AccountController) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	accounts, err := controller.service.FindAll(r.Context())
	if err != nil {
		response.CreateErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.CreateSuccessResponse(w, http.StatusOK, accounts)
}

func (controller AccountController) GetBalance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var accountId = params.ByName("account_id")

	account, err := controller.service.GetBalance(r.Context(), accountId)
	if err != nil {
		response.CreateErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.CreateSuccessResponse(w, http.StatusOK, account)
}
