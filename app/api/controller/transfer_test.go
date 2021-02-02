package controller

import (
	"bytes"
	"encoding/json"
	"github.com/brunodecastro/digital-accounts/app/api/auth"
	"github.com/brunodecastro/digital-accounts/app/api/response"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/custom-errors"
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/service"
	"github.com/brunodecastro/digital-accounts/app/util"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var (
	urlApiTransfer string
)

func init() {
	// Initialize app configs
	config.LoadConfigs()
}

func TestTransferController_Create_Success(t *testing.T) {
	t.Parallel()

	var transferInputVO = input.CreateTransferInputVO{
		AccountOriginId:      "0001",
		AccountDestinationId: "0002",
		Amount:               100,
	}
	body, _ := json.Marshal(transferInputVO)

	var transferOutputVO = output.CreateTransferOutputVO{
		Id:                   "0001",
		AccountOriginID:      "0001",
		AccountDestinationID: "0002",
		Amount:               types.Money(100).GetFloat64(),
		CreatedAt:            util.FormatDate(time.Time{}),
	}
	controller := NewTransferController(service.MockTransferService{
		ResultCreateTransfer: transferOutputVO,
		Err:                  nil,
	})

	wantStatusCode := http.StatusCreated
	wantResult := transferOutputVO

	endPoint := "/transfers"
	req, _ := http.NewRequest(http.MethodPost, endPoint, bytes.NewReader(body))
	rec := mockRequestHandler(req, http.MethodPost, endPoint, true, auth.AuthorizeMiddleware(controller.Create))

	// Check the response status code
	if statusCode := rec.Code; statusCode != wantStatusCode {
		t.Errorf("Create() error = %v, wantErr %v", wantStatusCode, statusCode)
	}

	// Check result response
	var responseResult output.CreateTransferOutputVO
	_ = json.Unmarshal(rec.Body.Bytes(), &responseResult)
	if !reflect.DeepEqual(responseResult, wantResult) {
		t.Errorf("Create() got = %v, wantResult %v", responseResult, wantResult)
	}
}

func TestTransferController_Create_Error_Validation(t *testing.T) {
	t.Parallel()

	var transferInputVO = input.CreateTransferInputVO{
		AccountOriginId:      "0001",
		AccountDestinationId: "0002",
		Amount:               0,
	}
	body, _ := json.Marshal(transferInputVO)

	var transferOutputVO = output.CreateTransferOutputVO{}
	controller := NewTransferController(service.MockTransferService{
		ResultCreateTransfer: transferOutputVO,
		Err:                  custom_errors.ErrorTransferAmountValue,
	})

	wantStatusCode := http.StatusBadRequest
	wantErr := true
	wantErrorMsg := custom_errors.ErrorTransferAmountValue.Error()

	endPoint := "/transfers"
	req, _ := http.NewRequest(http.MethodPost, endPoint, bytes.NewReader(body))
	rec := mockRequestHandler(req, http.MethodPost, endPoint, true, auth.AuthorizeMiddleware(controller.Create))

	// Check the response status code
	if statusCode := rec.Code; statusCode != wantStatusCode {
		t.Errorf("Create() error = %v, wantErr %v", wantStatusCode, statusCode)
	}

	// Check result response
	var responseError response.HttpErrorResponse
	_ = json.Unmarshal(rec.Body.Bytes(), &responseError)

	if wantErr && responseError.Error.Message != wantErrorMsg {
		t.Errorf("Create() error = %v, wantErr %v", responseError.Error.Message, wantErrorMsg)
		return
	}
}

func TestTransferController_Create_Error_Service(t *testing.T) {
	t.Parallel()

	var transferInputVO = input.CreateTransferInputVO{
		AccountOriginId:      "0001",
		AccountDestinationId: "0002",
		Amount:               100,
	}
	body, _ := json.Marshal(transferInputVO)

	var transferOutputVO = output.CreateTransferOutputVO{}
	controller := NewTransferController(service.MockTransferService{
		ResultCreateTransfer: transferOutputVO,
		Err:                  custom_errors.ErrorCreateAccount,
	})

	wantStatusCode := http.StatusInternalServerError
	wantErr := true
	wantErrorMsg := custom_errors.ErrorCreateAccount.Error()

	endPoint := "/transfers"
	req, _ := http.NewRequest(http.MethodPost, endPoint, bytes.NewReader(body))
	rec := mockRequestHandler(req, http.MethodPost, endPoint, true, auth.AuthorizeMiddleware(controller.Create))

	// Check the response status code
	if statusCode := rec.Code; statusCode != wantStatusCode {
		t.Errorf("Create() error = %v, wantErr %v", wantStatusCode, statusCode)
	}

	// Check result response
	var responseError response.HttpErrorResponse
	_ = json.Unmarshal(rec.Body.Bytes(), &responseError)

	if wantErr && responseError.Error.Message != wantErrorMsg {
		t.Errorf("Create() error = %v, wantErr %v", responseError.Error.Message, wantErrorMsg)
		return
	}
}

func TestTransferController_FindAll_Success(t *testing.T) {
	t.Parallel()

	var transfersOutputVO = []output.FindAllTransferOutputVO{
		{
			Id:                   "0001",
			AccountOriginID:      "0001",
			AccountDestinationID: "0002",
			Amount:               types.Money(100).GetFloat64(),
			CreatedAt:            util.FormatDate(time.Time{}),
		},
		{
			Id:                   "0002",
			AccountOriginID:      "0003",
			AccountDestinationID: "0003",
			Amount:               types.Money(150).GetFloat64(),
			CreatedAt:            util.FormatDate(time.Time{}),
		},
	}
	controller := NewTransferController(service.MockTransferService{
		ResultFindAll: transfersOutputVO,
		Err:           nil,
	})

	wantStatusCode := http.StatusOK
	wantResult := transfersOutputVO

	endPoint := "/transfers"
	req, _ := http.NewRequest(http.MethodGet, endPoint, nil)
	rec := mockRequestHandler(req, http.MethodGet, endPoint, true, auth.AuthorizeMiddleware(controller.FindAll))

	// Check the response status code
	if statusCode := rec.Code; statusCode != wantStatusCode {
		t.Errorf("Create() error = %v, wantErr %v", wantStatusCode, statusCode)
	}

	// Check result response
	var responseResult []output.FindAllTransferOutputVO
	_ = json.Unmarshal(rec.Body.Bytes(), &responseResult)
	if !reflect.DeepEqual(responseResult, wantResult) {
		t.Errorf("Create() got = %v, wantResult %v", responseResult, wantResult)
	}
}

func TestTransferController_FindAll_Error(t *testing.T) {
	t.Parallel()

	var transfersOutputVO []output.FindAllTransferOutputVO
	controller := NewTransferController(service.MockTransferService{
		ResultFindAll: transfersOutputVO,
		Err:           custom_errors.ErrorListingAllTransfers,
	})

	wantStatusCode := http.StatusInternalServerError
	wantErr := true
	wantErrorMsg := custom_errors.ErrorListingAllTransfers.Error()

	endPoint := "/transfers"
	req, _ := http.NewRequest(http.MethodGet, endPoint, nil)
	rec := mockRequestHandler(req, http.MethodGet, endPoint, true, auth.AuthorizeMiddleware(controller.FindAll))

	// Check the response status code
	if statusCode := rec.Code; statusCode != wantStatusCode {
		t.Errorf("Create() error = %v, wantErr %v", wantStatusCode, statusCode)
	}

	// Check result response
	var responseError response.HttpErrorResponse
	_ = json.Unmarshal(rec.Body.Bytes(), &responseError)

	if wantErr && responseError.Error.Message != wantErrorMsg {
		t.Errorf("Create() error = %v, wantErr %v", responseError.Error.Message, wantErrorMsg)
		return
	}
}
