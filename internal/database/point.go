package database

import (
	"time"

	"github.com/google/uuid"
)

type Point struct {
	ID        uuid.UUID `bun:"type:uuid,default:uuid_generate_v4(),pk"`
	EntityID  uuid.UUID `bun:"type:uuid,unique:entity_id_game_id_idx"`
	GameID    uuid.UUID `bun:"type:uuid,unique:entity_id_game_id_idx"`
	Count     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
