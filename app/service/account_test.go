package service

import (
	"context"
	"errors"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/custom-errors"
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/persistence/database/postgres"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/brunodecastro/digital-accounts/app/util"
	"reflect"
	"testing"
	"time"
)

var (
	transactionHelperMock postgres.MockTransactionHelper
)

func init() {
	// Initialize app configs
	config.LoadConfigs()

	transactionHelperMock = postgres.MockTransactionHelper{
		Result: context.Background(),
		Err:    nil,
	}
}

func Test_accountServiceImpl_Create(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository        repository.AccountRepository
		transactionHelper postgres.TransactionHelper
	}
	type args struct {
		ctx            context.Context
		accountInputVO input.CreateAccountInputVO
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		want         output.CreateAccountOutputVO
		wantErr      bool
		wantErrorMsg string
	}{

		{
			name: "Create Account test service Success",
			fields: fields{
				repository: repository.MockAccountRepositoryImpl{
					Result: model.Account{
						ID:        "0001",
						Cpf:       "00801246156",
						Name:      "Bruno 1",
						Secret:    "65O6G91K651",
						Balance:   0,
						CreatedAt: time.Time{},
					},
					ResultFindByCpf: nil,
					Err:             nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx: context.Background(),
				accountInputVO: input.CreateAccountInputVO{
					Cpf:     "008.012.461-56",
					Name:    "Bruno 1",
					Secret:  "65O6G91K651",
					Balance: 0,
				},
			},
			want: output.CreateAccountOutputVO{
				Cpf:       util.FormatCpf("00801246156"),
				Name:      "Bruno 1",
				Balance:   0,
				CreatedAt: util.FormatDate(time.Time{}),
			},
			wantErr: false,
		},
		{
			name: "Create Account test service Error - repository error",
			fields: fields{
				repository: repository.MockAccountRepositoryImpl{
					Result:          model.Account{},
					Err:             custom_errors.ErrorCreateAccount,
					ResultFindByCpf: nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx: context.Background(),
				accountInputVO: input.CreateAccountInputVO{
					Cpf:     "008.012.461-99",
					Name:    "Bruno 1",
					Secret:  "65O6G91K651",
					Balance: 0,
				},
			},
			want:         output.CreateAccountOutputVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorCreateAccount.Error(),
		},
		{
			name: "Create Account with the same cpf error",
			fields: fields{
				repository: repository.MockAccountRepositoryImpl{
					Result: model.Account{
						ID:        "0001",
						Cpf:       "00801246156",
						Name:      "Bruno 1",
						Secret:    "65O6G91K651",
						Balance:   0,
						CreatedAt: time.Time{},
					},
					ResultFindByCpf: &model.Account{
						ID:        "0001",
						Cpf:       "00801246156",
						Name:      "Bruno 1",
						Secret:    "65O6G91K651",
						Balance:   0,
						CreatedAt: time.Time{},
					},
					Err: nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx: context.Background(),
				accountInputVO: input.CreateAccountInputVO{
					Cpf:     "008.012.461-56",
					Name:    "Bruno 1",
					Secret:  "65O6G91K651",
					Balance: 0,
				},
			},
			want:         output.CreateAccountOutputVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorAccountCpfExists.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceImpl := NewAccountService(tt.fields.repository, tt.fields.transactionHelper)
			got, err := serviceImpl.Create(tt.args.ctx, tt.args.accountInputVO)
			if tt.wantErr && err.Error() != tt.wantErrorMsg {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountServiceImpl_GetAll(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository        repository.AccountRepository
		transactionHelper postgres.TransactionHelper
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		want         []output.FindAllAccountOutputVO
		wantErr      bool
		wantErrorMsg string
	}{
		{
			name: "Get all accounts test service success",
			fields: fields{
				repository: repository.MockAccountRepositoryImpl{
					Results: []model.Account{
						{
							ID:        "0001",
							Cpf:       "00801246156",
							Name:      "Bruno 1",
							Secret:    "65O6G91K651",
							Balance:   100,
							CreatedAt: time.Time{},
						},
						{
							ID:        "0002",
							Cpf:       "00801246157",
							Name:      "Bruno 2",
							Secret:    "65O6G91K6510",
							Balance:   250,
							CreatedAt: time.Time{},
						},
					},
					Err: nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx: context.Background(),
			},
			want: []output.FindAllAccountOutputVO{
				{
					ID:        "0001",
					Cpf:       util.FormatCpf("00801246156"),
					Name:      "Bruno 1",
					Balance:   types.Money(100).ToFloat64(),
					CreatedAt: util.FormatDate(time.Time{}),
				},
				{
					ID:        "0002",
					Cpf:       util.FormatCpf("00801246157"),
					Name:      "Bruno 2",
					Balance:   types.Money(250).ToFloat64(),
					CreatedAt: util.FormatDate(time.Time{}),
				},
			},
			wantErr: false,
		},
		{
			name: "Get all accounts test service error - repository error",
			fields: fields{
				repository: repository.MockAccountRepositoryImpl{
					Results: []model.Account{},
					Err:     errors.New("error on get all accounts"),
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx: context.Background(),
			},
			want:         []output.FindAllAccountOutputVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorListingAllAccounts.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceImpl := accountServiceImpl{
				repository: tt.fields.repository,
			}
			got, err := serviceImpl.FindAll(tt.args.ctx)
			if tt.wantErr && err.Error() != tt.wantErrorMsg {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountServiceImpl_GetBalance(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository        repository.AccountRepository
		transactionHelper postgres.TransactionHelper
	}
	type args struct {
		ctx       context.Context
		accountID string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		want         output.FindAccountBalanceOutputVO
		wantErr      bool
		wantErrorMsg string
	}{
		{
			name: "Get balance test service success",
			fields: fields{
				repository: repository.MockAccountRepositoryImpl{
					Result: model.Account{
						ID:        "0001",
						Cpf:       "00801246156",
						Name:      "Bruno 1",
						Secret:    "65O6G91K651",
						Balance:   100,
						CreatedAt: time.Time{},
					},
					Err: nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx:       context.Background(),
				accountID: "0001",
			},
			want: output.FindAccountBalanceOutputVO{
				ID:      "0001",
				Balance: types.Money(100).ToFloat64(),
			},
			wantErr: false,
		},
		{
			name: "Get balance test service - repository error",
			fields: fields{
				repository: repository.MockAccountRepositoryImpl{
					Result: model.Account{},
					Err:    errors.New("error get balance"),
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx:       context.Background(),
				accountID: "0001",
			},
			want:         output.FindAccountBalanceOutputVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorGettingAccountBalance.Error(),
		},
		{
			name: "Get balance test service error - account not found",
			fields: fields{
				repository: repository.MockAccountRepositoryImpl{
					Result: model.Account{},
					Err:    nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx:       context.Background(),
				accountID: "0001",
			},
			want:         output.FindAccountBalanceOutputVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorAccountNotFound.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceImpl := accountServiceImpl{
				repository: tt.fields.repository,
			}
			got, err := serviceImpl.GetBalance(tt.args.ctx, tt.args.accountID)
			if tt.wantErr && err.Error() != tt.wantErrorMsg {
				t.Errorf("GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}
