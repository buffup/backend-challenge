package database

import (
	"time"

	"github.com/google/uuid"
)

type Leaderboard struct {
	ID        uuid.UUID
	GameID    []uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
