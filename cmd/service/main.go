package main

import (
	"context"

	"github.com/buffup/backend-challenge/internal/database"
)

func main() {
	dsn := "postgres://postgres:postgres@localhost:55021/points?sslmode=disable"

	db, err := database.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := database.Migrate(context.Background(), dsn); err != nil {
		panic(err)
	}

}
