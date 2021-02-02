package postgres

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/constants"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// TransactionHelper - transaction helper interface
type TransactionHelper interface {
	StartTransaction(ctx context.Context) (context.Context, error)
	CommitTransaction(ctx context.Context) error
	RollbackTransaction(ctx context.Context) error
	GetTransactionFromContext(ctx context.Context) pgx.Tx
}

type transactionHelperImpl struct {
	dataBasePool *pgxpool.Pool
}

// NewTransactionHelper - transaction helper instance
func NewTransactionHelper(dataBasePool *pgxpool.Pool) TransactionHelper {
	return &transactionHelperImpl{
		dataBasePool: dataBasePool,
	}
}

// StartTransaction - starts a transaction
func (transaction transactionHelperImpl) StartTransaction(ctx context.Context) (context.Context, error) {
	tx, err := transaction.dataBasePool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	newContext := setTransactionInContext(ctx, tx)
	return newContext, nil
}

// CommitTransaction - commits a transaction
func (transaction transactionHelperImpl) CommitTransaction(ctx context.Context) error {
	tx := transaction.GetTransactionFromContext(ctx)
	err := tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

// RollbackTransaction - rollbacks a transaction
func (transaction transactionHelperImpl) RollbackTransaction(ctx context.Context) error {
	tx := transaction.GetTransactionFromContext(ctx)
	err := tx.Rollback(ctx)
	if err != nil {
		return err
	}
	return nil
}

// setTransactionInContext - set the transaction number in the context
func setTransactionInContext(ctx context.Context, tx pgx.Tx) context.Context {
	newContext := context.WithValue(ctx, constants.TransactionContextKey, tx)
	return newContext
}

// GetTransactionFromContext - gets the transaction number in the context
func (transaction transactionHelperImpl) GetTransactionFromContext(ctx context.Context) pgx.Tx {
	return ctx.Value(constants.TransactionContextKey).(pgx.Tx)
}
