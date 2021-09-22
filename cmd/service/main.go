package main

import (
	"context"

	"github.com/buffup/backend-challenge/internal/database"
)

func main() {
	db, err := database.Connect(context.Background())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := database.CreateTables(context.Background(), db); err != nil {
		panic(err)
	}

}
