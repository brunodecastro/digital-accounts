package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/persistence/database/postgres"
	"github.com/brunodecastro/digital-accounts/app/service"
	"github.com/brunodecastro/digital-accounts/app/util"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

var (
	urlAPI                string
	transactionHelperMock postgres.MockTransactionHelper
)

func init() {
	// Initialize app configs
	apiConfig := config.GetAPIConfigs()

	urlAPI = apiConfig.WebServerConfig.GetWebServerAddress()

	transactionHelperMock = postgres.MockTransactionHelper{
		Result: context.Background(),
		Err:    nil,
	}
}

func TestAccountController_Create(t *testing.T) {
	t.Parallel()

	endPoint := fmt.Sprintf("%s/accounts", urlAPI)

	accountInputVOTest := input.CreateAccountInputVO{
		CPF:     "008.012.461-56",
		Name:    "Bruno de Castro Oliveira",
		Secret:  "123456",
		Balance: 1050,
	}
	body, _ := json.Marshal(accountInputVOTest)

	rec := httptest.NewRecorder()
	resp := httptest.NewRequest(http.MethodPost, endPoint, bytes.NewReader(body))

	accountOutputVO := output.CreateAccountOutputVO{
		CPF:       util.FormatCpf(accountInputVOTest.CPF),
		Name:      accountInputVOTest.Name,
		Balance:   types.Money(accountInputVOTest.Balance).ToFloat64(),
		CreatedAt: util.FormatDate(time.Time{}),
	}

	type fields struct {
		service           service.AccountService
		transactionHelper postgres.TransactionHelper
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
		_ httprouter.Params
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
		wantResult     output.CreateAccountOutputVO
		wantErr        bool
	}{
		{
			name: "Create account controller success",
			fields: fields{
				service: service.MockAccountService{
					ResultCreateAccount: accountOutputVO,
					Err:                 nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				w: rec,
				r: resp,
			},
			wantStatusCode: http.StatusCreated,
			wantResult:     accountOutputVO,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := AccountController{
				service: tt.fields.service,
			}

			controller.Create(tt.args.w, tt.args.r, nil)

			// Check the response status code
			if statusCode := rec.Code; statusCode != tt.wantStatusCode {
				t.Errorf("Create() error = %v, wantErr %v", tt.wantStatusCode, statusCode)
			}

			// Check result response
			var responseResult = output.CreateAccountOutputVO{}
			_ = json.Unmarshal(rec.Body.Bytes(), &responseResult)

			if !reflect.DeepEqual(responseResult, tt.wantResult) {
				t.Errorf("Create() got = %v, wantResult %v", responseResult, tt.wantResult)
			}
		})
	}
}

func TestAccountController_GetAll(t *testing.T) {
	t.Parallel()

	endPoint := fmt.Sprintf("%s/accounts", urlAPI)
	rec := httptest.NewRecorder()
	resp := httptest.NewRequest(http.MethodGet, endPoint, nil)

	accountsOutputVO := []output.FindAllAccountOutputVO{
		{
			ID:        "0001",
			Cpf:       "008.012.461-56",
			Name:      "Bruno 1",
			Balance:   15,
			CreatedAt: util.FormatDate(time.Time{}),
		},
		{
			ID:        "0002",
			Cpf:       "00801246157",
			Name:      "Bruno 2",
			Balance:   25.5,
			CreatedAt: util.FormatDate(time.Time{}),
		},
	}

	type fields struct {
		service           service.AccountService
		transactionHelper postgres.TransactionHelper
	}
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
		wantResult     []output.FindAllAccountOutputVO
		wantErr        bool
	}{
		{
			name: "Create account controller success",
			fields: fields{
				service: service.MockAccountService{
					ResultGetAll: accountsOutputVO,
					Err:          nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				w: rec,
				r: resp,
			},
			wantStatusCode: http.StatusOK,
			wantResult:     accountsOutputVO,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := AccountController{
				service: tt.fields.service,
			}

			controller.FindAll(tt.args.w, tt.args.r, nil)

			// Check the response status code
			if statusCode := rec.Code; statusCode != tt.wantStatusCode {
				t.Errorf("Create() error = %v, wantErr %v", tt.wantStatusCode, statusCode)
			}

			// Check result response
			var responseResult []output.FindAllAccountOutputVO
			_ = json.Unmarshal(rec.Body.Bytes(), &responseResult)
			if !reflect.DeepEqual(responseResult, tt.wantResult) {
				t.Errorf("Create() got = %v, wantResult %v", responseResult, tt.wantResult)
			}
		})
	}
}

func TestAccountController_GetBalance(t *testing.T) {
	t.Parallel()

	var accountID = "0001"
	endPoint := fmt.Sprintf("%s/account/%s/balance", urlAPI, accountID)
	rec := httptest.NewRecorder()
	resp := httptest.NewRequest(http.MethodGet, endPoint, nil)

	accountOutputVO := output.FindAccountBalanceOutputVO{
		ID:      accountID,
		Balance: types.Money(250).ToFloat64(),
	}

	type fields struct {
		service           service.AccountService
		transactionHelper postgres.TransactionHelper
	}
	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		params httprouter.Params
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
		wantResult     output.FindAccountBalanceOutputVO
		wantErr        bool
	}{
		{
			name: "Get account balance controller success",
			fields: fields{
				service: service.MockAccountService{
					ResultGetBalance: accountOutputVO,
					Err:              nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				w: rec,
				r: resp,
			},
			wantStatusCode: http.StatusOK,
			wantResult:     accountOutputVO,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := NewAccountController(tt.fields.service)

			params := []httprouter.Param{
				{Key: "account_id", Value: accountID},
			}
			controller.GetBalance(tt.args.w, tt.args.r, params)

			// Check the response status code
			if statusCode := rec.Code; statusCode != tt.wantStatusCode {
				t.Errorf("Create() error = %v, wantErr %v", tt.wantStatusCode, statusCode)
			}

			// Check result response
			var responseResult output.FindAccountBalanceOutputVO
			_ = json.Unmarshal(rec.Body.Bytes(), &responseResult)
			if !reflect.DeepEqual(responseResult, tt.wantResult) {
				t.Errorf("Create() got = %v, wantResult %v", responseResult, tt.wantResult)
			}
		})
	}
}
