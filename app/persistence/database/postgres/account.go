package postgres

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type accountRepositoryImpl struct {
	dataBasePool      *pgxpool.Pool
	transactionHelper TransactionHelper
}

func NewAccountRepository(dataBasePool *pgxpool.Pool, transactionHelper TransactionHelper) repository.AccountRepository {
	return &accountRepositoryImpl{
		dataBasePool:      dataBasePool,
		transactionHelper: transactionHelper,
	}
}

func (repositoryImpl accountRepositoryImpl) Create(ctx context.Context, account model.Account) (*model.Account, error) {
	tx := repositoryImpl.transactionHelper.GetTransactionFromContext(ctx)

	var sqlQuery = `
		INSERT INTO 
			accounts (id, name, cpf, secret, balance, created_at)
		VALUES 
			($1, $2, $3, $4, $5, $6)
	`

	_, err := tx.Exec(
		ctx,
		sqlQuery,
		account.Id,
		account.Name,
		account.Cpf,
		account.Secret,
		account.Balance,
		account.CreatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error creating account")
	}

	return &account, nil
}

func (repositoryImpl accountRepositoryImpl) UpdateBalance(ctx context.Context, accountOriginId model.AccountID, balance types.Money) (int64, error) {
	tx := repositoryImpl.transactionHelper.GetTransactionFromContext(ctx)

	var sqlQuery = `
		UPDATE accounts
			SET balance = $1
		WHERE
			id = $2
	`

	result, err := tx.Exec(
		ctx,
		sqlQuery,
		balance,
		accountOriginId,
	)
	if err != nil {
		return 0, errors.Wrap(err, "error updating account balance")
	}

	return result.RowsAffected(), nil
}

func (repositoryImpl accountRepositoryImpl) FindAll(ctx context.Context) ([]model.Account, error) {
	var sqlQuery = `
		SELECT
			id, name, cpf, secret, balance, created_at
		FROM
			accounts
	`

	rows, err := repositoryImpl.dataBasePool.Query(ctx, sqlQuery)
	if err != nil {
		return []model.Account{}, errors.Wrap(err, "error executing listing all accounts query")
	}
	defer rows.Close()

	var accounts = make([]model.Account, 0)
	for rows.Next() {
		var account = model.Account{}

		err = rows.Scan(
			&account.Id,
			&account.Name,
			&account.Cpf,
			&account.Secret,
			&account.Balance,
			&account.CreatedAt)
		if err != nil {
			return []model.Account{}, errors.Wrap(err, "error scanning listing all accounts")
		}
		accounts = append(accounts, account)
	}

	if err = rows.Err(); err != nil {
		return []model.Account{}, err
	}

	return accounts, nil
}

func (repositoryImpl accountRepositoryImpl) GetBalance(ctx context.Context, accountId string) (*model.Account, error) {
	var sqlQuery = `
		SELECT
			id, balance
		FROM
			accounts
		where
			id = $1
	`
	var account = model.Account{}

	row := repositoryImpl.dataBasePool.QueryRow(ctx, sqlQuery, accountId)
	err := row.Scan(&account.Id, &account.Balance)

	if err != nil && err != pgx.ErrNoRows {
		return nil, errors.Wrap(err, "error scanning accounts balance")
	}

	return &account, nil
}

func (repositoryImpl accountRepositoryImpl) FindByCpf(ctx context.Context, cpf string) (*model.Account, error) {
	var sqlQuery = `
		SELECT
			id, name, cpf, secret, balance, created_at
		FROM
			accounts
		where
			cpf = $1
	`
	var account = model.Account{}

	row := repositoryImpl.dataBasePool.QueryRow(ctx, sqlQuery, cpf)
	err := row.Scan(
		&account.Id,
		&account.Name,
		&account.Cpf,
		&account.Secret,
		&account.Balance,
		&account.CreatedAt)

	if err != nil && err != pgx.ErrNoRows {
		return nil, errors.Wrap(err, "error scanning find account by cpf")
	}
	return &account, nil
}

func (repositoryImpl accountRepositoryImpl) FindById(ctx context.Context, accountId string) (*model.Account, error) {
	var sqlQuery = `
		SELECT
			id, name, cpf, secret, balance, created_at
		FROM
			accounts
		where
			id = $1
	`
	var account = model.Account{}

	row := repositoryImpl.dataBasePool.QueryRow(ctx, sqlQuery, accountId)
	err := row.Scan(
		&account.Id,
		&account.Name,
		&account.Cpf,
		&account.Secret,
		&account.Balance,
		&account.CreatedAt)

	if err != nil && err != pgx.ErrNoRows {
		return nil, errors.Wrap(err, "error scanning find account by cpf")
	}
	return &account, nil
}
