package postgres

import (
	"database/sql"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/util"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func UpMigrations(databasePostgresConfig *config.DatabasePostgresConfig) error {
	databaseConnection, err := sql.Open("postgres", databasePostgresConfig.GetDatabaseURL());
	util.MaybeFatal(err,"Unable to open postgres connection to run migrations")

	databaseDriver, err := postgres.WithInstance(databaseConnection, &postgres.Config{})
	defer databaseDriver.Close()

	if err != nil {
		return err
	}

	migrateInstance, err := migrate.NewWithDatabaseInstance("file://migrations", databasePostgresConfig.DatabaseName, databaseDriver)
	defer migrateInstance.Close()

	if err != nil {
		return err
	}

	err = migrateInstance.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
