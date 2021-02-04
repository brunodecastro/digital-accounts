package postgres

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/persistence/database/postgres/fakes"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"reflect"
	"sync"
	"testing"
)

var (
	dataBasePool      *pgxpool.Pool
	transactionHelper TransactionHelper
	doOnce            sync.Once
)

// TestMain test dockertest database setup
func TestMain(m *testing.M) {
	var dockerCleanFunc func() error
	var err error
	pgxPool, dockerCleanFunc, err := StartDockerTestPostgresDataBase()
	if err != nil {
		log.Fatalf("error setting up docker db: %v", err)
	}

	// setup data repository
	setupDataRepository(pgxPool)

	// setup fake accounts in the database
	SetupFakeAccounts(pgxPool)

	// run tests
	runCode := m.Run()

	// close db connection
	pgxPool.Close()

	// cleanup docker container
	err = dockerCleanFunc()
	if err != nil {
		log.Fatalf("error cleaning up docker container: %v", err)
	}

	os.Exit(runCode)
}

func Test_accountRepositoryImpl_FindAll(t *testing.T) {

	fakeAccounts := fakes.GetFakeAccounts()

	type fields struct {
		dataBasePool      *pgxpool.Pool
		transactionHelper TransactionHelper
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Account
		wantErr bool
	}{
		{
			name: "find all accounts persistence - success",
			fields: fields{
				dataBasePool:      dataBasePool,
				transactionHelper: transactionHelper,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    *fakeAccounts,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repositoryImpl := NewAccountRepository(tt.fields.dataBasePool, tt.fields.transactionHelper)

			got, err := repositoryImpl.FindAll(tt.args.ctx)
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

func Test_accountRepositoryImpl_FindByCPF(t *testing.T) {

	accountFake1 := fakes.GetFakeAccount1()

	type fields struct {
		dataBasePool      *pgxpool.Pool
		transactionHelper TransactionHelper
	}
	type args struct {
		ctx context.Context
		cpf string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Account
		wantErr bool
	}{
		{
			name: "find by cpf test - success",
			fields: fields{
				dataBasePool:      dataBasePool,
				transactionHelper: transactionHelper,
			},
			args: args{
				ctx: context.Background(),
				cpf: accountFake1.CPF,
			},
			want:    &accountFake1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repositoryImpl := NewAccountRepository(tt.fields.dataBasePool, tt.fields.transactionHelper)

			got, err := repositoryImpl.FindByCPF(tt.args.ctx, tt.args.cpf)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByCPF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByCPF() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountRepositoryImpl_FindByID(t *testing.T) {
	accountFake1 := fakes.GetFakeAccount1()

	type fields struct {
		dataBasePool      *pgxpool.Pool
		transactionHelper TransactionHelper
	}
	type args struct {
		ctx       context.Context
		accountID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Account
		wantErr bool
	}{
		{
			name: "find account by id test - success",
			fields: fields{
				dataBasePool:      dataBasePool,
				transactionHelper: transactionHelper,
			},
			args: args{
				ctx:       context.Background(),
				accountID: string(accountFake1.ID),
			},
			want:    &accountFake1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repositoryImpl := NewAccountRepository(tt.fields.dataBasePool, tt.fields.transactionHelper)

			got, err := repositoryImpl.FindByID(tt.args.ctx, tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountRepositoryImpl_GetBalance(t *testing.T) {
	accountFake1 := fakes.GetFakeAccount1()

	type fields struct {
		dataBasePool      *pgxpool.Pool
		transactionHelper TransactionHelper
	}
	type args struct {
		ctx       context.Context
		accountID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "find account balance - success",
			fields: fields{
				dataBasePool:      dataBasePool,
				transactionHelper: transactionHelper,
			},
			args: args{
				ctx:       context.Background(),
				accountID: string(accountFake1.ID),
			},
			want:    accountFake1.Balance.ToInt64(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repositoryImpl := NewAccountRepository(tt.fields.dataBasePool, tt.fields.transactionHelper)

			got, err := repositoryImpl.GetBalance(tt.args.ctx, tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Balance.ToInt64() != tt.want {
				t.Errorf("GetBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accountRepositoryImpl_UpdateBalance(t *testing.T) {

	accountFake2 := fakes.GetFakeAccount2()

	type fields struct {
		dataBasePool      *pgxpool.Pool
		transactionHelper TransactionHelper
	}
	type args struct {
		ctx       context.Context
		accountID model.AccountID
		balance   types.Money
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "update account balance - success",
			fields: fields{
				dataBasePool:      dataBasePool,
				transactionHelper: transactionHelper,
			},
			args: args{
				ctx:       context.Background(),
				accountID: accountFake2.ID,
				balance:   900,
			},
			want:    900,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repositoryImpl := NewAccountRepository(tt.fields.dataBasePool, tt.fields.transactionHelper)

			transactionContext, err := tt.fields.transactionHelper.StartTransaction(tt.args.ctx)
			if err != nil {
				log.Fatal(err.Error())
			}

			rowsAfeccted, err := repositoryImpl.UpdateBalance(transactionContext, tt.args.accountID, tt.args.balance)

			if err != nil || rowsAfeccted < 1 {
				t.Errorf("Error on update account balance")
				tt.fields.transactionHelper.RollbackTransaction(transactionContext)
			}
			tt.fields.transactionHelper.CommitTransaction(transactionContext)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			accountBalanceUpdated, err := repositoryImpl.FindByID(context.Background(), string(tt.args.accountID))

			if err != nil && (accountBalanceUpdated == nil || accountBalanceUpdated.Balance.ToInt64() != tt.want) {
				t.Errorf("UpdateBalance() got = %v, want %v", accountBalanceUpdated.Balance.ToInt64(), tt.want)
			}
		})
	}
}

func Test_accountRepositoryImpl_Create(t *testing.T) {

	accountFake := fakes.GenerateNewFakeAccount("Fake Created")
	type fields struct {
		dataBasePool      *pgxpool.Pool
		transactionHelper TransactionHelper
	}
	type args struct {
		ctx     context.Context
		account model.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Account
		wantErr bool
	}{
		{
			name: "create account persistence - success",
			fields: fields{
				dataBasePool:      dataBasePool,
				transactionHelper: transactionHelper,
			},
			args: args{
				ctx:     context.Background(),
				account: accountFake,
			},
			want:    &accountFake,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			repositoryImpl := NewAccountRepository(tt.fields.dataBasePool, tt.fields.transactionHelper)

			transactionContext, err := tt.fields.transactionHelper.StartTransaction(tt.args.ctx)
			if err != nil {
				log.Fatal(err.Error())
			}

			accountCreated, err := repositoryImpl.Create(transactionContext, tt.args.account)
			if err != nil {
				t.Errorf("Error on create account")
				tt.fields.transactionHelper.RollbackTransaction(transactionContext)
			}
			tt.fields.transactionHelper.CommitTransaction(transactionContext)

			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}

			accountFindByID, err := repositoryImpl.FindByID(context.Background(), string(accountCreated.ID))
			if !reflect.DeepEqual(*accountCreated, *accountFindByID) {
				t.Errorf("Create() got = %v, want %v", *accountCreated, *accountFindByID)
			}
		})
	}
}

func setupDataRepository(pgxPool *pgxpool.Pool) {
	doOnce.Do(func() {
		dataBasePool = pgxPool
		transactionHelper = NewTransactionHelper(pgxPool)
	})
}
