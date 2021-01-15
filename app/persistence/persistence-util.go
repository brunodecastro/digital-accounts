package persistence

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/logger"
	"github.com/brunodecastro/digital-accounts/app/util/conditional"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
)

// PoolConfig gets the pool configuration from the database.
func PoolConfig(databaseConfig *config.DatabaseConfig) *pgxpool.Pool {
	databaseConnectionConfig, err := pgxpool.ParseConfig(databaseConfig.GetDatabaseDSN())
	conditional.MaybeFatal(err, "Unable to parse the pool config")

	databaseConnectionConfig.ConnConfig.Logger = zapadapter.NewLogger(logger.LogApp.GetZapLogImplementation())

	databaseConnection, err := pgxpool.ConnectConfig(context.Background(), databaseConnectionConfig)
	conditional.MaybeFatal(err, "Unable to connect the pool config")

	return databaseConnection
}
