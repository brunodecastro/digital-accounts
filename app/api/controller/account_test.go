package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	apiConfig          *config.Config
	accountInputVOTest input.CreateAccountInputVO
	urlApi             string
)

func init() {
	// Initialize app configs
	apiConfig = config.LoadConfigs()

	urlApi = apiConfig.WebServerConfig.GetWebServerAddress()

	accountInputVOTest = input.CreateAccountInputVO{
		Cpf:     "008.012.461-56",
		Name:    "Bruno de Castro Oliveira",
		Secret:  "123456",
		Balance: 1050,
	}
}

func TestAccountController_Create(t *testing.T) {
	t.Parallel()

	body, _ := json.Marshal(accountInputVOTest)
	endPoint := fmt.Sprintf("%s/accounts", urlApi)

	type fields struct {
		service service.AccountService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
		_ httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Create account controller success",
			fields: fields{
				service: service.MockAccountService{
					ResultCreateAccount: output.CreateAccountOutputVO{},
					Err:                 nil,
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, endPoint, bytes.NewReader(body)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := AccountController{
				service: tt.fields.service,
			}
			controller.Create(tt.args.w, tt.args.r, nil)

		})
	}
}

func TestAccountController_GetAll(t *testing.T) {
	type fields struct {
		service service.AccountService
	}
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		in2 httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := AccountController{
				service: tt.fields.service,
			}
			fmt.Print(controller)
		})
	}
}

func TestAccountController_GetBalance(t *testing.T) {
	type fields struct {
		service service.AccountService
	}
	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		params httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := AccountController{
				service: tt.fields.service,
			}
			fmt.Print(controller)
		})
	}
}
