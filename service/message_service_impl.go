package service

import (
	"context"
	"database/sql"

	"github.com/artamananda/artanymous/helper"
	"github.com/artamananda/artanymous/model/domain"
	"github.com/artamananda/artanymous/model/web"
	"github.com/artamananda/artanymous/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type MessageServiceImpl struct {
	MessageRepository repository.MessageRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func (service *MessageServiceImpl) Create(ctx context.Context, request web.MessageCreateRequest) web.MessageResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	message := domain.Message{
		Question: request.Question,
	}

	message = service.MessageRepository.Save(ctx, tx, message)

	return helper.ToMessageResponse(message)
}

func (service *MessageServiceImpl) Delete(ctx context.Context, messageId uuid.UUID) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	message, err := service.MessageRepository.FindById(ctx, tx, messageId)
	helper.PanicIfError(err)

	service.MessageRepository.Delete(ctx, tx, message)
}

func (service *MessageServiceImpl) FindById(ctx context.Context, messageId uuid.UUID) web.MessageResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	message, err := service.MessageRepository.FindById(ctx, tx, messageId)
	helper.PanicIfError(err)

	return helper.ToMessageResponse(message)
}

func (service *MessageServiceImpl) FindAll(ctx context.Context) []web.MessageResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	messages := service.MessageRepository.FindAll(ctx, tx)

	return helper.ToMessageResponses(messages)
}
