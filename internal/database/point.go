package database

import (
	"time"

	"github.com/google/uuid"
)

type Point struct {
	ID        uuid.UUID
	EntityID  uuid.UUID
	GameID    uuid.UUID
	Count     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
