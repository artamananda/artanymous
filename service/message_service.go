package service

import (
	"context"

	"github.com/artamananda/artanymous/model/web"
	"github.com/google/uuid"
)

type MessageService interface {
	Create(ctx context.Context, request web.MessageCreateRequest) web.MessageResponse
	Delete(ctx context.Context, messageId uuid.UUID)
	FindById(ctx context.Context, messageId uuid.UUID) web.MessageResponse
	FindAll(ctx context.Context) []web.MessageResponse
}
