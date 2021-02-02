package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
)

// MockTransactionHelper - struct that represents transaction helper
type MockTransactionHelper struct {
	Result context.Context
	Err    error
}

// StartTransaction - starts a transaction mock
func (m MockTransactionHelper) StartTransaction(ctx context.Context) (context.Context, error) {
	return m.Result, m.Err
}

// CommitTransaction - commits a transaction mock
func (m MockTransactionHelper) CommitTransaction(ctx context.Context) error {
	return m.Err
}

// RollbackTransaction - rollbacks a transaction mock
func (m MockTransactionHelper) RollbackTransaction(ctx context.Context) error {
	return m.Err
}

// GetTransactionFromContext - gets the transaction number in the context mock
func (m MockTransactionHelper) GetTransactionFromContext(ctx context.Context) pgx.Tx {
	return nil
}
