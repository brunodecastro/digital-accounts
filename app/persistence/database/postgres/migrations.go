package postgres

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"       // driver to get migrations file
)

// UpMigrations - up the database migrations
func UpMigrations(postgresURI string, migrationsPath string) error {

	migrateInstance, err := migrate.New(fmt.Sprintf("file://%s", migrationsPath), postgresURI)
	defer migrateInstance.Close()

	if err != nil {
		fmt.Errorf("error on migration instance: %v", err)
		return err
	}

	err = migrateInstance.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
