package postgres

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/logger"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/util"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
)

// ConnectPoolConfig - gets the pool configuration from the database.
func ConnectPoolConfig() *pgxpool.Pool {
	var databaseConnectionString string

	var databaseURLConfig = config.GetAPIConfigs().DatabaseConfig.DatabaseURL
	if  databaseURLConfig != "" {
		// if DatabaseURL is set, then returns the complete url
		databaseConnectionString = databaseURLConfig
	} else {
		databaseConnectionString = config.GetAPIConfigs().DatabaseConfig.GetDatabaseDSN()
	}
	
	databaseConnectionConfig, err := pgxpool.ParseConfig(databaseConnectionString)
	util.MaybeFatal(err, "Unable to parse the pool config")

	databaseConnectionConfig.ConnConfig.Logger = zapadapter.NewLogger(logger.GetLogger())

	databaseConnection, err := pgxpool.ConnectConfig(context.Background(), databaseConnectionConfig)
	util.MaybeFatal(err, "Unable to connect the pool config")

	return databaseConnection
}
