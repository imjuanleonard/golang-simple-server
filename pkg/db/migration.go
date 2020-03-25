package db

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/imjuanleonard/palu-covid/config"
)

const databaseMigrationPath = "file://./db/migrations"

func Migrate() error {
	m, err := getMigration(config.Database.ConnectionURL)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return err
		}
	}

	return nil
}

func Rollback() error {
	m, err := getMigration(config.Database.ConnectionURL)
	if err != nil {
		return err
	}

	if err := m.Down(); err != nil {
		if err != migrate.ErrNoChange {
			return err
		}
	}

	return nil
}

func getMigration(connection string) (*migrate.Migrate, error) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(databaseMigrationPath, "postgres", driver)
	if err != nil {
		return nil, err
	}

	return m, nil
}
