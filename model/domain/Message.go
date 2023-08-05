package domain

import (
	"time"
)

type Message struct {
	Id        int
	Question  string
	CreatedAt time.Time
}
