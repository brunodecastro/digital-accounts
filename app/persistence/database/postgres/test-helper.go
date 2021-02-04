package postgres

import (
	"context"
	"fmt"
	"github.com/brunodecastro/digital-accounts/app/persistence/database/postgres/fakes"
	"github.com/brunodecastro/digital-accounts/app/util"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/pkg/errors"
	"time"
)

// StartDockerTestPostgresDataBase is a helper method to automatically start a postgres docker instance via dockertest.
func StartDockerTestPostgresDataBase() (*pgxpool.Pool, func() error, error) {
	var pgxPool *pgxpool.Pool
	var postgresURI string

	// create dockertest pool
	dockerPool, err := dockertest.NewPool("")
	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating new pool in dockertest database")
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := dockerPool.Run("postgres", "latest", []string{"POSTGRES_PASSWORD=secret"})
	if err != nil {
		return nil, nil, errors.Wrap(err, "error creating resource dockertest database")
	}

	// 10 second wait time to connect to the db
	dockerPool.MaxWait = time.Second * 10
	err = dockerPool.Retry(func() error {
		postgresURI = fmt.Sprintf(
			"postgres://postgres:secret@localhost:%s/postgres?sslmode=disable",
			resource.GetPort("5432/tcp"),
		)
		pgxPool, err = pgxpool.Connect(context.Background(), postgresURI)
		if err != nil {
			return err
		}
		conn, err := pgxPool.Acquire(context.Background())
		if err != nil {
			return err
		}
		defer conn.Release()
		return conn.Conn().Ping(context.Background())
	})
	if err != nil {
		return nil, nil, errors.Wrap(err, "error connecting to dockertest database")
	}

	// run migrations
	err = UpMigrations(postgresURI, "migrations")
	util.MaybeFatal(err, "Unable to execute postgres migrations.")

	return pgxPool, func() error { return dockerPool.Purge(resource) }, nil
}

// SetupFakeAccounts - inserts fake accounts in the database to test with dockertest
func SetupFakeAccounts(pgxPool *pgxpool.Pool) error {
	transactionHelper := NewTransactionHelper(pgxPool)
	repositoryImpl := NewAccountRepository(pgxPool, transactionHelper)

	for _, accountFake := range *fakes.GetFakeAccounts() {

		transactionContext, err := transactionHelper.StartTransaction(context.Background())
		if err != nil {
			return errors.Wrap(err, "error on start transaction")
		}

		_, err = repositoryImpl.Create(transactionContext, accountFake)
		if err != nil {
			transactionHelper.RollbackTransaction(transactionContext)
			return errors.Wrap(err, "error on create fake account")
		}
		transactionHelper.CommitTransaction(transactionContext)
		if err != nil {
			return errors.Wrap(err, "error on commit transaction")
		}
	}
	return nil
}
