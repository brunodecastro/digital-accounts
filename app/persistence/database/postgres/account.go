package postgres

import (
	"context"
	"database/sql"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type accountRepositoryImpl struct {
	dataBasePool *pgxpool.Pool
}

func NewAccountRepository(dataBasePool *pgxpool.Pool) repository.AccountRepository {
	return &accountRepositoryImpl{
		dataBasePool: dataBasePool,
	}
}

func (repositoryImpl accountRepositoryImpl) Create(ctx context.Context, account model.Account) (*model.Account, error) {
	tx, err := repositoryImpl.dataBasePool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	var sqlQuery = `
		INSERT INTO 
			accounts (id, name, cpf, secret, balance, created_at)
		VALUES 
			($1, $2, $3, $4, $5, $6)
	`

	_, err = tx.Exec(
		ctx,
		sqlQuery,
		account.ID,
		account.Name,
		account.Cpf,
		account.Secret,
		account.Balance,
		account.CreatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "error creating account")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (repositoryImpl accountRepositoryImpl) GetAll(ctx context.Context) ([]model.Account, error) {
	var sqlQuery = `
		SELECT id, name, cpf, secret, balance, created_at FROM accounts
	`

	rows, err := repositoryImpl.dataBasePool.Query(ctx, sqlQuery)
	if err != nil {
		return []model.Account{}, errors.Wrap(err, "error listing all accounts")
	}
	defer rows.Close()

	var accounts = make([]model.Account, 0)
	for rows.Next() {
		var account = model.Account{}

		if err = rows.Scan(&account.ID, &account.Name, &account.Cpf, &account.Secret, &account.Balance, &account.CreatedAt); err != nil {
			return []model.Account{}, errors.Wrap(err, "error listing all accounts")
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
		SELECT id, balance FROM accounts where id = $1
	`
	var account = model.Account{}

	row := repositoryImpl.dataBasePool.QueryRow(ctx, sqlQuery, accountId)
	err := row.Scan(&account.ID, &account.Balance)

	if err != nil && err != sql.ErrNoRows {
		return nil, nil
	}

	return &account, nil
}