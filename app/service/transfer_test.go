package service

import (
	"context"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/errors"
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

func init() {
	// Initialize app configs
	config.GetAPIConfigs()
}

func Test_transferServiceImpl_Create(t *testing.T) {
	t.Parallel()

	transactionHelperMock := postgres.MockTransactionHelper{
		Result: context.Background(),
		Err:    nil,
	}

	type fields struct {
		transferRepository repository.TransferRepository
		accountRepository  repository.AccountRepository
		transactionHelper  postgres.TransactionHelper
	}
	type args struct {
		ctx             context.Context
		transferInputVO input.CreateTransferInputVO
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		want         output.CreateTransferOutputVO
		wantErr      bool
		wantErrorMsg string
	}{
		{
			name: "Create transfer success",
			fields: fields{
				transferRepository: repository.MockTransferRepositoryImpl{
					Result: model.Transfer{
						ID:                   "0001",
						AccountOriginID:      "0001",
						AccountDestinationID: "0002",
						Amount:               10,
						CreatedAt:            time.Time{},
					},
					Err: nil,
				},
				accountRepository: repository.MockAccountRepositoryImpl{
					ResultFindByID: &model.Account{
						ID:        "0002",
						CPF:       "00801246157",
						Name:      "Bruno 2",
						Secret:    "65O6G91K651",
						Balance:   1000,
						CreatedAt: time.Time{},
					},
					ResultBalance: 1,
					Err:           nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx: context.Background(),
				transferInputVO: input.CreateTransferInputVO{
					AccountOriginID:      "0001",
					AccountDestinationID: "0002",
					Amount:               10,
				},
			},
			want: output.CreateTransferOutputVO{
				ID:                   "0001",
				AccountOriginID:      "0001",
				AccountDestinationID: "0002",
				Amount:               types.Money(10).ToFloat64(),
				CreatedAt:            util.FormatDate(time.Time{}),
			},
			wantErr:      false,
			wantErrorMsg: "",
		},
		{
			name: "Create transfer - repository create error",
			fields: fields{
				transferRepository: repository.MockTransferRepositoryImpl{
					Result: model.Transfer{
						ID:                   "0001",
						AccountOriginID:      "0001",
						AccountDestinationID: "0002",
						Amount:               10,
						CreatedAt:            time.Time{},
					},
					Err: custom_errors.ErrorCreateTransfer,
				},
				accountRepository: repository.MockAccountRepositoryImpl{
					ResultFindByID: &model.Account{
						ID:        "0002",
						CPF:       "00801246157",
						Name:      "Bruno 2",
						Secret:    "65O6G91K651",
						Balance:   1000,
						CreatedAt: time.Time{},
					},
					ResultBalance: 1,
					Err:           nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx: context.Background(),
				transferInputVO: input.CreateTransferInputVO{
					AccountOriginID:      "0001",
					AccountDestinationID: "0002",
					Amount:               10,
				},
			},
			want:         output.CreateTransferOutputVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorCreateTransfer.Error(),
		},
		{
			name: "Create transfer - account origin not found",
			fields: fields{
				transferRepository: repository.MockTransferRepositoryImpl{
					Result: model.Transfer{
						ID:                   "0001",
						AccountOriginID:      "0001",
						AccountDestinationID: "0002",
						Amount:               10,
						CreatedAt:            time.Time{},
					},
					Err: nil,
				},
				accountRepository: repository.MockAccountRepositoryImpl{
					ResultFindByID: nil,
					Err:            nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx: context.Background(),
				transferInputVO: input.CreateTransferInputVO{
					AccountOriginID:      "0001",
					AccountDestinationID: "0002",
					Amount:               10,
				},
			},
			want:         output.CreateTransferOutputVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorAccountOriginNotFound.Error(),
		},
		{
			name: "Create transfer - amount <= 0",
			fields: fields{
				transferRepository: repository.MockTransferRepositoryImpl{
					Result: model.Transfer{
						ID:                   "0001",
						AccountOriginID:      "0001",
						AccountDestinationID: "0002",
						Amount:               10,
						CreatedAt:            time.Time{},
					},
					Err: nil,
				},
				accountRepository: repository.MockAccountRepositoryImpl{
					ResultFindByID: nil,
					Err:            nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx: context.Background(),
				transferInputVO: input.CreateTransferInputVO{
					AccountOriginID:      "0001",
					AccountDestinationID: "0002",
					Amount:               0,
				},
			},
			want:         output.CreateTransferOutputVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorTransferAmountValue.Error(),
		},
		{
			name: "Create transfer - insufficient balance",
			fields: fields{
				transferRepository: repository.MockTransferRepositoryImpl{
					Result: model.Transfer{
						ID:                   "0001",
						AccountOriginID:      "0001",
						AccountDestinationID: "0002",
						Amount:               10000,
						CreatedAt:            time.Time{},
					},
					Err: nil,
				},
				accountRepository: repository.MockAccountRepositoryImpl{
					ResultFindByID: &model.Account{
						ID:        "0002",
						CPF:       "00801246157",
						Name:      "Bruno 2",
						Secret:    "65O6G91K651",
						Balance:   10,
						CreatedAt: time.Time{},
					},
					Err: nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx: context.Background(),
				transferInputVO: input.CreateTransferInputVO{
					AccountOriginID:      "0001",
					AccountDestinationID: "0002",
					Amount:               1000,
				},
			},
			want:         output.CreateTransferOutputVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorInsufficientBalance.Error(),
		},
		{
			name: "Create transfer - amount <= 0",
			fields: fields{
				transferRepository: repository.MockTransferRepositoryImpl{
					Result: model.Transfer{
						ID:                   "0001",
						AccountOriginID:      "0001",
						AccountDestinationID: "0002",
						Amount:               10,
						CreatedAt:            time.Time{},
					},
					Err: nil,
				},
				accountRepository: repository.MockAccountRepositoryImpl{
					ResultFindByID: nil,
					Err:            nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx: context.Background(),
				transferInputVO: input.CreateTransferInputVO{
					AccountOriginID:      "0001",
					AccountDestinationID: "0002",
					Amount:               0,
				},
			},
			want:         output.CreateTransferOutputVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorTransferAmountValue.Error(),
		},
		{
			name: "Create transfer error - transfer to the same account",
			fields: fields{
				transferRepository: repository.MockTransferRepositoryImpl{
					Result: model.Transfer{},
					Err:    nil,
				},
				accountRepository: repository.MockAccountRepositoryImpl{
					Err: nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx: context.Background(),
				transferInputVO: input.CreateTransferInputVO{
					AccountOriginID:      "0001",
					AccountDestinationID: "0001",
					Amount:               10,
				},
			},
			want:         output.CreateTransferOutputVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorTransferSameAccount.Error(),
		},
		{
			name: "Create transfer - update balance error",
			fields: fields{
				transferRepository: repository.MockTransferRepositoryImpl{
					Result: model.Transfer{
						ID:                   "0001",
						AccountOriginID:      "0001",
						AccountDestinationID: "0002",
						Amount:               10,
						CreatedAt:            time.Time{},
					},
					Err: nil,
				},
				accountRepository: repository.MockAccountRepositoryImpl{
					ResultFindByID: &model.Account{
						ID:        "0002",
						CPF:       "00801246157",
						Name:      "Bruno 2",
						Secret:    "65O6G91K651",
						Balance:   1000,
						CreatedAt: time.Time{},
					},
					ResultBalance: 0,
					Err:           nil,
				},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx: context.Background(),
				transferInputVO: input.CreateTransferInputVO{
					AccountOriginID:      "0001",
					AccountDestinationID: "0002",
					Amount:               10,
				},
			},
			want:         output.CreateTransferOutputVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorUpdateAccountOriginBalance.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceImpl := NewTransferService(
				tt.fields.transferRepository,
				tt.fields.accountRepository,
				tt.fields.transactionHelper,
			)
			got, err := serviceImpl.Create(tt.args.ctx, tt.args.transferInputVO)
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

func Test_transferServiceImpl_FindAll(t *testing.T) {
	t.Parallel()

	transactionHelperMock := postgres.MockTransactionHelper{
		Result: context.Background(),
		Err:    nil,
	}

	type fields struct {
		transferRepository repository.TransferRepository
		accountRepository  repository.AccountRepository
		transactionHelper  postgres.TransactionHelper
	}
	type args struct {
		ctx       context.Context
		accountID string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		want         []output.FindAllTransferOutputVO
		wantErr      bool
		wantErrorMsg string
	}{
		{
			name: "Find all transfers success",
			fields: fields{
				transferRepository: repository.MockTransferRepositoryImpl{
					Results: []model.Transfer{
						{
							ID:                   "0001",
							AccountOriginID:      "0001",
							AccountDestinationID: "0002",
							Amount:               10,
							CreatedAt:            time.Time{},
						},
						{
							ID:                   "0002",
							AccountOriginID:      "0001",
							AccountDestinationID: "0002",
							Amount:               15,
							CreatedAt:            time.Time{},
						},
					},
					Err: nil,
				},
				accountRepository: repository.MockAccountRepositoryImpl{},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx: context.Background(),
			},
			want: []output.FindAllTransferOutputVO{
				{
					ID:                   "0001",
					AccountOriginID:      "0001",
					AccountDestinationID: "0002",
					Amount:               types.Money(10).ToFloat64(),
					CreatedAt:            util.FormatDate(time.Time{}),
				},
				{
					ID:                   "0002",
					AccountOriginID:      "0001",
					AccountDestinationID: "0002",
					Amount:               types.Money(15).ToFloat64(),
					CreatedAt:            util.FormatDate(time.Time{}),
				},
			},
			wantErr:      false,
			wantErrorMsg: "",
		},
		{
			name: "Find all transfers error",
			fields: fields{
				transferRepository: repository.MockTransferRepositoryImpl{
					Results: []model.Transfer{},
					Err:     custom_errors.ErrorListingAllTransfers,
				},
				accountRepository: repository.MockAccountRepositoryImpl{},
				transactionHelper: transactionHelperMock,
			},
			args: args{
				ctx:       context.Background(),
				accountID: "0001",
			},
			want:         []output.FindAllTransferOutputVO{},
			wantErr:      true,
			wantErrorMsg: custom_errors.ErrorListingAllTransfers.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceImpl := transferServiceImpl{
				transferRepository: tt.fields.transferRepository,
				accountRepository:  tt.fields.accountRepository,
				transactionHelper:  tt.fields.transactionHelper,
			}
			got, err := serviceImpl.FindAll(tt.args.ctx, tt.args.accountID)
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
