package database

import (
	"context"
	"fmt"
	"net/http"

	"github.com/buffup/backend-challenge/migrations"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
)

func Migrate(ctx context.Context, dsn string) error {
	source, err := httpfs.New(http.FS(migrations.FS), ".")
	if err != nil {
		return fmt.Errorf("failed to load migrations: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("httpfs", source, dsn)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	return m.Up()
}
