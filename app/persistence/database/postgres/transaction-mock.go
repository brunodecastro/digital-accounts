package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type MockTransactionHelper struct {
	Result context.Context
	Err    error
}

func (m MockTransactionHelper) StartTransaction(ctx context.Context) (context.Context, error) {
	return m.Result, m.Err
}

func (m MockTransactionHelper) CommitTransaction(ctx context.Context) error {
	return m.Err
}

func (m MockTransactionHelper) RollbackTransaction(ctx context.Context) error {
	return m.Err
}

func (m MockTransactionHelper) GetTransactionFromContext(ctx context.Context) pgx.Tx {
	return nil
}
