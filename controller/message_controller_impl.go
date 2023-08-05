package controller

import (
	"net/http"

	"github.com/artamananda/artanymous/helper"
	"github.com/artamananda/artanymous/model/web"
	"github.com/artamananda/artanymous/service"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type MessageControllerImpl struct {
	MessageService service.MessageService
}

func NewMessageController(messageService service.MessageService) MessageController {
	return &MessageControllerImpl{
		MessageService: messageService,
	}
}

func (controller *MessageControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	messageCreateRequest := web.MessageCreateRequest{}
	helper.ReadFromRequestBody(request, &messageCreateRequest)

	messageResponse := controller.MessageService.Create(request.Context(), messageCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   messageResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MessageControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	messageId := params.ByName("messageId")
	id, err := uuid.Parse(messageId)
	helper.PanicIfError(err)

	controller.MessageService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MessageControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	messageId := params.ByName("messageId")
	id, err := uuid.Parse(messageId)
	helper.PanicIfError(err)

	messageResponse := controller.MessageService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   messageResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MessageControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	messageResponses := controller.MessageService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   messageResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
