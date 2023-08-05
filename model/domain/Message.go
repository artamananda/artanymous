package domain

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id        uuid.UUID
	Question  string
	CreatedAt time.Time
}
