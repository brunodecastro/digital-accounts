package postgres

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/logger"
	"github.com/brunodecastro/digital-accounts/app/util"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
)

// PoolConfig gets the pool configuration from the database.
func ConnectPoolConfig(databaseConfig *config.DatabasePostgresConfig) *pgxpool.Pool {
	databaseConnectionConfig, err := pgxpool.ParseConfig(databaseConfig.GetDatabaseDSN())
	util.MaybeFatal(err, "Unable to parse the pool config")

	databaseConnectionConfig.ConnConfig.Logger = zapadapter.NewLogger(logger.LogApp.GetZapLogImplementation())

	databaseConnection, err := pgxpool.ConnectConfig(context.Background(), databaseConnectionConfig)
	util.MaybeFatal(err, "Unable to connect the pool config")

	return databaseConnection
}
