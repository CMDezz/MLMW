package utils

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Automatically migrate database when server up
// if there is no change, it wont update the dtb
func MigrateDatabase(migrationDir string, dbSource string) error {
	migration, err := migrate.New(migrationDir, dbSource)
	if err != nil {
		return err
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
