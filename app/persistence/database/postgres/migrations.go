package postgres

import (
	"database/sql"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/util"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // driver to get migrations file
	"log"
)

// UpMigrations - up the database migrations
func UpMigrations(databasePostgresConfig *config.DatabasePostgresConfig) error {
	databaseConnection, err := sql.Open("postgres", databasePostgresConfig.GetDatabaseURL())
	util.MaybeFatal(err, "Unable to open postgres connection to run migrations")

	databaseDriver, err := postgres.WithInstance(databaseConnection, &postgres.Config{})
	defer databaseDriver.Close()

	if err != nil {
		return err
	}

	migrateInstance, err := migrate.NewWithDatabaseInstance("file://app/persistence/database/postgres/migrations", databasePostgresConfig.DatabaseName, databaseDriver)
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

func UpMigrations2(pgURI string) error {
	migrateInstance, err := migrate.New("file://migrations", pgURI)
	defer migrateInstance.Close()
	if err != nil && err.Error() != "no change" {
		log.Fatalf("couldn't run migrations: %v", err)
		return err
	}

	err = migrateInstance.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("couldn't up migrations: %v", err)
		return err
	}

	return nil
}

