package repository

import (
	"context"
	"database/sql"

	"github.com/artamananda/artanymous/model/domain"
)

type MessageRepository interface {
	Save(ctx context.Context, tx *sql.Tx, message domain.Message) domain.Message
	Delete(ctx context.Context, tx *sql.Tx, message domain.Message) domain.Message
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Message
}
