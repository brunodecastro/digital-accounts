package controller

import (
	"encoding/json"
	"github.com/brunodecastro/digital-accounts/app/api/response"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TransferController struct {
	service service.TransferService
}

func NewTransferController(service service.TransferService) TransferController {
	return TransferController{
		service: service,
	}
}

func (controller TransferController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	var transferInputVO input.CreateTransferInputVO
	if err := json.NewDecoder(r.Body).Decode(&transferInputVO); err != nil {
		response.CreateErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// TODO: validate fields

	transferCreated, err := controller.service.Create(r.Context(), transferInputVO)
	if err != nil {
		response.CreateErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.CreateSuccessResponse(w, http.StatusCreated, transferCreated)
}

func (controller TransferController) FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	transfers, err := controller.service.FindAll(r.Context())
	if err != nil {
		response.CreateErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.CreateSuccessResponse(w, http.StatusOK, transfers)
}
