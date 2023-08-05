package web

import (
	"time"

	"github.com/google/uuid"
)

type MessageResponse struct {
	Id        uuid.UUID
	Question  string
	CreatedAt time.Time
}
