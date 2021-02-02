package service

import (
	"context"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/custom-errors"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/brunodecastro/digital-accounts/app/util"
	"reflect"
	"testing"
	"time"
)

func init() {
	// Initialize app configs
	config.LoadConfigs()
}

func Test_authenticationServiceImpl_Authenticate(t *testing.T) {
	t.Parallel()

	type fields struct {
		accountRepository repository.AccountRepository
	}
	type args struct {
		ctx               context.Context
		credentialInputVO input.CredentialInputVO
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		want         vo.CredentialVO
		wantErr      bool
		wantErrorMsg string
	}{
		{
			name: "Authentication Success",
			fields: fields{
				accountRepository: repository.MockAccountRepositoryImpl{
					ResultFindByCpf: &model.Account{
						ID:        "0001",
						Cpf:       "00801246156",
						Name:      "Bruno 1",
						Secret:    util.EncryptPassword("65O6G91K651"),
						Balance:   0,
						CreatedAt: time.Time{},
					},
					Err: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				credentialInputVO: input.CredentialInputVO{
					Cpf:    "008.012.461-56",
					Secret: "65O6G91K651",
				},
			},
			want: vo.CredentialVO{
				Cpf:       "00801246156",
				AccountID: "0001",
				Username:  "Bruno 1",
			},
			wantErr:      false,
			wantErrorMsg: "",
		},
		{
			name: "Authentication Failed - invalid cpf",
			fields: fields{
				accountRepository: repository.MockAccountRepositoryImpl{
					ResultFindByCpf: nil,
					Err:             nil,
				},
			},
			args: args{
				ctx: context.Background(),
				credentialInputVO: input.CredentialInputVO{
					Cpf:    "008.012.461-59",
					Secret: "65O6G91K651",
				},
			},
			want:         vo.CredentialVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorCpfInvalid.Error(),
		},
		{
			name: "Authentication Failed - wrong password",
			fields: fields{
				accountRepository: repository.MockAccountRepositoryImpl{
					ResultFindByCpf: &model.Account{
						ID:        "0001",
						Cpf:       "00801246156",
						Name:      "Bruno 1",
						Secret:    util.EncryptPassword("65O6G91K651"),
						Balance:   0,
						CreatedAt: time.Time{},
					},
					Err: nil,
				},
			},
			args: args{
				ctx: context.Background(),
				credentialInputVO: input.CredentialInputVO{
					Cpf:    "008.012.461-56",
					Secret: "wrong-password",
				},
			},
			want:         vo.CredentialVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorCredentialWrongSecret.Error(),
		},
		{
			name: "Authentication Failed - account not found",
			fields: fields{
				accountRepository: repository.MockAccountRepositoryImpl{
					ResultFindByCpf: nil,
					Err:             nil,
				},
			},
			args: args{
				ctx: context.Background(),
				credentialInputVO: input.CredentialInputVO{
					Cpf:    "008.012.461-56",
					Secret: "65O6G91K651",
				},
			},
			want:         vo.CredentialVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorAccountNotFound.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceImpl := NewAuthenticationService(tt.fields.accountRepository)
			got, err := serviceImpl.Authenticate(tt.args.ctx, tt.args.credentialInputVO)
			if tt.wantErr && err.Error() != tt.wantErrorMsg {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authenticate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
