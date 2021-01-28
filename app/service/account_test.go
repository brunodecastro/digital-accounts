package service

import (
	"context"
	"errors"
	"github.com/brunodecastro/digital-accounts/app/common"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/brunodecastro/digital-accounts/app/util"
	"reflect"
	"testing"
	"time"
)

func Test_accountServiceImpl_Create(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository repository.AccountRepository
	}
	type args struct {
		ctx            context.Context
		accountInputVO input.CreateAccountInputVO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    output.CreateAccountOutputVO
		wantErr bool
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
					Err: nil,
				},
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
			name: "Create Account test service Error",
			fields: fields{
				repository: repository.MockAccountRepositoryImpl{
					Result: model.Account{},
					Err:    errors.New("error on create account"),
				},
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
			want:    output.CreateAccountOutputVO{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceImpl := accountServiceImpl{
				repository: tt.fields.repository,
			}
			got, err := serviceImpl.Create(tt.args.ctx, tt.args.accountInputVO)
			if (err != nil) != tt.wantErr {
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
		repository repository.AccountRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []output.FindAllAccountOutputVO
		wantErr bool
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
			},
			args: args{
				ctx: context.Background(),
			},
			want: []output.FindAllAccountOutputVO{
				{
					Id:        "0001",
					Cpf:       util.FormatCpf("00801246156"),
					Name:      "Bruno 1",
					Balance:   common.Money(100).GetFloat(),
					CreatedAt: util.FormatDate(time.Time{}),
				},
				{
					Id:        "0002",
					Cpf:       util.FormatCpf("00801246157"),
					Name:      "Bruno 2",
					Balance:   common.Money(250).GetFloat(),
					CreatedAt: util.FormatDate(time.Time{}),
				},
			},
			wantErr: false,
		},
		{
			name: "Get all accounts test service error",
			fields: fields{
				repository: repository.MockAccountRepositoryImpl{
					Results: []model.Account{},
					Err:     errors.New("error on get all accounts"),
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want:    []output.FindAllAccountOutputVO{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceImpl := accountServiceImpl{
				repository: tt.fields.repository,
			}
			got, err := serviceImpl.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountServiceImpl_GetBalance(t *testing.T) {
	t.Parallel()

	type fields struct {
		repository repository.AccountRepository
	}
	type args struct {
		ctx       context.Context
		accountId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    output.FindAccountBalanceOutputVO
		wantErr bool
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
			},
			args: args{
				ctx:       context.Background(),
				accountId: "0001",
			},
			want: output.FindAccountBalanceOutputVO{
				Id:      "0001",
				Balance: common.Money(100).GetFloat(),
			},
			wantErr: false,
		},
		{
			name: "Get balance test service error",
			fields: fields{
				repository: repository.MockAccountRepositoryImpl{
					Result: model.Account{},
					Err:    errors.New("error get balance"),
				},
			},
			args: args{
				ctx:       context.Background(),
				accountId: "0001",
			},
			want:    output.FindAccountBalanceOutputVO{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceImpl := accountServiceImpl{
				repository: tt.fields.repository,
			}
			got, err := serviceImpl.GetBalance(tt.args.ctx, tt.args.accountId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}
