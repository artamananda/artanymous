package web

import (
	"time"

	"github.com/google/uuid"
)

type MessageResponse struct {
	Id        uuid.UUID `json:"id"`
	Question  string    `json:"question"`
	CreatedAt time.Time `json:"created_at"`
}
