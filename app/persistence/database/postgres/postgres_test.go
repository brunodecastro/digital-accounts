package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
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
	// setup fake transfers in the database
	SetupFakeTransfers(pgxPool)

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

func setupDataRepository(pgxPool *pgxpool.Pool) {
	doOnce.Do(func() {
		dataBasePool = pgxPool
		transactionHelper = NewTransactionHelper(pgxPool)
	})
}
