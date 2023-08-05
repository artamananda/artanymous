package repository

import (
	"context"
	"database/sql"

	"github.com/artamananda/artanymous/model/domain"
	"github.com/google/uuid"
)

type MessageRepository interface {
	Save(ctx context.Context, tx *sql.Tx, message domain.Message) domain.Message
	Delete(ctx context.Context, tx *sql.Tx, message domain.Message)
	FindById(ctx context.Context, tx *sql.Tx, messageId uuid.UUID) (domain.Message, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Message
}
