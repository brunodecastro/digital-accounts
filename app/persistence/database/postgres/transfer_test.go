package postgres

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/jackc/pgx/v4/pgxpool"
	"reflect"
	"testing"
)




func Test_transferRepositoryImpl_FindAll(t *testing.T) {
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
			name: "",
			fields:  fields{
				dataBasePool:      nil,
				transactionHelper: nil,
			},
			args:    args{
				ctx:             nil,
				accountOriginID: "",
			},
			want:    nil,
			wantErr: false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repositoryImpl := transferRepositoryImpl{
				dataBasePool:      tt.fields.dataBasePool,
				transactionHelper: tt.fields.transactionHelper,
			}
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repositoryImpl := transferRepositoryImpl{
				dataBasePool:      tt.fields.dataBasePool,
				transactionHelper: tt.fields.transactionHelper,
			}
			got, err := repositoryImpl.Create(tt.args.ctx, tt.args.transfer)
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
