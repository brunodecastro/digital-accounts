package postgres

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type transferRepositoryImpl struct {
	dataBasePool      *pgxpool.Pool
	transactionHelper TransactionHelper
}

// NewTransferRepository - transfer repository instance
func NewTransferRepository(dataBasePool *pgxpool.Pool, transactionHelper TransactionHelper) repository.TransferRepository {
	return &transferRepositoryImpl{
		dataBasePool:      dataBasePool,
		transactionHelper: transactionHelper,
	}
}

// Create - creates a new transfer
func (repositoryImpl transferRepositoryImpl) Create(ctx context.Context, transfer model.Transfer) (*model.Transfer, error) {
	tx := repositoryImpl.transactionHelper.GetTransactionFromContext(ctx)

	var sqlQuery = `
		INSERT INTO 
			transfers (id, account_origin_id, account_destination_id, amount, created_at)
		VALUES 
			($1, $2, $3, $4, $5)
	`

	_, err := tx.Exec(
		ctx,
		sqlQuery,
		transfer.ID,
		transfer.AccountOriginID,
		transfer.AccountDestinationID,
		transfer.Amount,
		transfer.CreatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error executing creating transfer query")
	}

	return &transfer, nil
}

// FindAll - list all transfers
func (repositoryImpl transferRepositoryImpl) FindAll(ctx context.Context, accountOriginId string) ([]model.Transfer, error) {
	var sqlQuery = `
		SELECT 
			id, account_origin_id, account_destination_id, amount, created_at 
		FROM 
			transfers
		where
			account_origin_id = $1
	`

	rows, err := repositoryImpl.dataBasePool.Query(ctx, sqlQuery, accountOriginId)
	if err != nil {
		return []model.Transfer{}, errors.Wrap(err, "error executing listing transfers query")
	}
	defer rows.Close()

	var transfers = make([]model.Transfer, 0)
	for rows.Next() {
		var transfer = model.Transfer{}

		err = rows.Scan(
			&transfer.ID,
			&transfer.AccountOriginID,
			&transfer.AccountDestinationID,
			&transfer.Amount,
			&transfer.CreatedAt)
		if err != nil {
			return []model.Transfer{}, errors.Wrap(err, "error scanning transfers")
		}
		transfers = append(transfers, transfer)
	}

	if err = rows.Err(); err != nil {
		return []model.Transfer{}, err
	}

	return transfers, nil
}
