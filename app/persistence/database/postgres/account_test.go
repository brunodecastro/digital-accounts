package postgres

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
	"reflect"
	"sync"
	"testing"
	"time"
)

var (
	dataBasePool      *pgxpool.Pool
	transactionHelper TransactionHelper
	//accountRepository repository.AccountRepository
	doOnce sync.Once
)

// TestMain test dockertest database setup
func TestMain(m *testing.M) {
	var dockerCleanFunc func() error
	var err error
	pgxPool, dockerCleanFunc, err := StartDockerTestPostgresDataBase()
	if err != nil {
		log.Fatalf("error setting up docker db: %v", err)
	}

	// setup repository
	setupRepository(pgxPool)

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

func TestAccountRepositoryImpl_Create(t *testing.T) {
	t.Parallel()

	accountInput := model.Account{
		ID:        model.AccountID(common.NewUUID()),
		CPF:       "00801246156",
		Name:      "Bruno de Castro Oliveira",
		Secret:    "123456",
		Balance:   1000,
		CreatedAt: time.Time{},
	}

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
		want    model.Account
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				ctx:     context.Background(),
				account: accountInput,
			},
			fields: fields{
				dataBasePool:      dataBasePool,
				transactionHelper: transactionHelper,
			},
			want:    accountInput,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			accountRepository := NewAccountRepository(tt.fields.dataBasePool, tt.fields.transactionHelper)

			accountCreated, err := accountRepository.Create(tt.args.ctx, tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(accountCreated, tt.want) {
				t.Errorf("Create() got = %v, want %v", accountCreated, tt.want)
			}
		})
	}
}

func setupRepository(pgxPool *pgxpool.Pool) {
	doOnce.Do(func() {
		dataBasePool = pgxPool
		transactionHelper = NewTransactionHelper(pgxPool)
		//accountRepository = NewAccountRepository(pgxPool, transactionHelper)
	})
}
