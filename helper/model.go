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

func ToMessageResponses(messages []domain.Message) []web.MessageResponse {
	var messageResponses []web.MessageResponse
	for _, message := range messages {
		messageResponses = append(messageResponses, ToMessageResponse(message))
	}
	return messageResponses
}
