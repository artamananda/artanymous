package repository

import (
	"context"
	"database/sql"

	"github.com/artamananda/artanymous/model/domain"
)

type MessageRepositoryImpl struct {
}

func (repository *MessageRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, message domain.Message) domain.Message {
	panic("err")
}

func (repository *MessageRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, message domain.Message) domain.Message {
	panic("err")
}

func (repository *MessageRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Message {
	panic("err")
}
