package postgres

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/persistence/database/postgres/fakes"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"reflect"
	"testing"
)

func Test_transferRepositoryImpl_FindAll(t *testing.T) {

	fakeTransfers := fakes.GetFakeTransfers()

	type fields struct {
		dataBasePool      *pgxpool.Pool
		transactionHelper TransactionHelper
	}
	type args struct {
		ctx             context.Context
		accountOriginID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Transfer
		wantErr bool
	}{
		{
			name: "find all transfers - success",
			fields: fields{
				dataBasePool:      dataBasePool,
				transactionHelper: transactionHelper,
			},
			args: args{
				ctx:             context.Background(),
				accountOriginID: string(fakes.GetFakeAccount1().ID),
			},
			want:    *fakeTransfers,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repositoryImpl := NewTransferRepository(tt.fields.dataBasePool, tt.fields.transactionHelper)

			got, err := repositoryImpl.FindAll(tt.args.ctx, tt.args.accountOriginID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transferRepositoryImpl_Create(t *testing.T) {
	transferFake := fakes.GenerateNewFakeTransfer()

	type fields struct {
		dataBasePool      *pgxpool.Pool
		transactionHelper TransactionHelper
	}
	type args struct {
		ctx      context.Context
		transfer model.Transfer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Transfer
		wantErr bool
	}{
		{
			name: "create transfer - success ",
			fields: fields{
				dataBasePool:      dataBasePool,
				transactionHelper: transactionHelper,
			},
			args: args{
				ctx:      context.Background(),
				transfer: transferFake,
			},
			want:    &transferFake,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repositoryImpl := NewTransferRepository(tt.fields.dataBasePool, tt.fields.transactionHelper)

			transactionContext, err := tt.fields.transactionHelper.StartTransaction(tt.args.ctx)
			if err != nil {
				log.Fatal(err.Error())
			}

			transferCreated, err := repositoryImpl.Create(transactionContext, tt.args.transfer)
			if err != nil {
				t.Errorf("Error on create transfer")
				tt.fields.transactionHelper.RollbackTransaction(transactionContext)
			}
			tt.fields.transactionHelper.CommitTransaction(transactionContext)

			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(*transferCreated, *tt.want) {
				t.Errorf("Create() got = %v, want %v", *transferCreated, *tt.want)
			}
		})
	}
}
