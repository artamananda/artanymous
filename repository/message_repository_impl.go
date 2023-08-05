package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/artamananda/artanymous/helper"
	"github.com/artamananda/artanymous/model/domain"
	"github.com/google/uuid"
)

type MessageRepositoryImpl struct {
}

func (repository *MessageRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, message domain.Message) domain.Message {
	id := uuid.New()
	SQL := "insert into messages(id, question, created_at) value (?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, id, message.Question, time.Now())
	helper.PanicIfError(err)

	message.Id = id
	message.CreatedAt = time.Now()
	return message
}

func (repository *MessageRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, message domain.Message) {
	SQL := "delete from messages where id=?"
	_, err := tx.ExecContext(ctx, SQL, message.Id)
	helper.PanicIfError(err)
}

func (repository *MessageRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Message {
	SQL := "select id, question, created_at from messages"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var messages []domain.Message
	for rows.Next() {
		message := domain.Message{}
		err := rows.Scan(&message.Id, &message.Question, &message.CreatedAt)
		helper.PanicIfError(err)
		messages = append(messages, message)
	}
	return messages
}