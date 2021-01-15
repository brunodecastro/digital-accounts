package persistence

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/error"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

// PoolConfig gets the pool configuration from the database.
func PoolConfig(databaseConfig *config.DatabaseConfig, log *logrus.Logger) *pgxpool.Pool {
	databaseConnectionConfig, err := pgxpool.ParseConfig(databaseConfig.GetDatabaseDSN())
	error.MaybeFatal(err, "Unable to parse the pool config")

	databaseConnectionConfig.ConnConfig.Logger = logrusadapter.NewLogger(log)

	databaseConnection, err := pgxpool.ConnectConfig(context.Background(), databaseConnectionConfig)
	error.MaybeFatal(err, "Unable to connect the pool config")

	return databaseConnection
}
