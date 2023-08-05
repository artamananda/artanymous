package helper

import (
	"github.com/artamananda/artanymous/model/domain"
	"github.com/artamananda/artanymous/model/web"
)

func ToMessageResponse(message domain.Message) web.MessageResponse {
	return web.MessageResponse{
		Id:        message.Id,
		Question:  message.Question,
		CreatedAt: message.CreatedAt,
	}
}
