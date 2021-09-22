package database

import (
	"time"

	"github.com/google/uuid"
)

type Leaderboard struct {
	ID        uuid.UUID   `bun:"type:uuid,default:uuid_generate_v4(),pk"`
	GameID    []uuid.UUID `bun:"type:uuid[],array"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
