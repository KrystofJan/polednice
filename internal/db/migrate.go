package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
)

func MigrateUp() error {
	connStr, err := GetConnString()
	if err != nil {
		return err
	}

	wd, _ := os.Getwd()
	migrationPath := "file://" + filepath.Join(wd, "internal", "migrations")
	migration, err := migrate.New(migrationPath, connStr)
	if err != nil {
		return fmt.Errorf("There was a problem with creating migrations: %v", err)
	}
	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("There was a problem with migrations: %v", err)
	}
	return nil
}
