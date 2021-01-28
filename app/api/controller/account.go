package controller

import (
	"encoding/json"
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

func (controller AccountController) Create(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	defer request.Body.Close()

	var accountInputVO input.CreateAccountInputVO
	if err := json.NewDecoder(request.Body).Decode(&accountInputVO); err != nil {
		// TODO: encapsular erros 400
		http.Error(response, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// TODO: validar campos

	accountCreated, err := controller.service.Create(request.Context(), accountInputVO)
	if err != nil {
		// TODO: encapsular erros 500
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(accountCreated)
}

func (controller AccountController) GetAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//return controller.service.GetAll(ctx)
}

func (controller AccountController) GetBalance(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//return controller.service.GetBalance(ctx, accountId)
}

func createJsonSuccessResponse(response http.ResponseWriter, statusCode int) {
	// TODO: criar uma resposta generica
}