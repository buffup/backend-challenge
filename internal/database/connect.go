package database

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// Connect establishes a connection to the provided database dsn using the standard library driver
// Once a client has been created run any necessary migrations and then return the ent client
func Connect(ctx context.Context) (*bun.DB, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/points?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	return db, nil
}
