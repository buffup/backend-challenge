package database

import (
	"context"

	"github.com/uptrace/bun"
)

func CreateTables(ctx context.Context, db *bun.DB) error {
	models := []interface{}{
		&Point{},
		&Leaderboard{},
	}
	for _, m := range models {
		_, err := db.NewCreateTable().Model(m).IfNotExists().Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
