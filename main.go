package main

import (
	"net/http"

	_ "github.com/lib/pq"

	"github.com/artamananda/artanymous/config"
	"github.com/artamananda/artanymous/controller"
	"github.com/artamananda/artanymous/helper"
	"github.com/artamananda/artanymous/repository"
	"github.com/artamananda/artanymous/service"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := config.NewDB()
	validate := validator.New()
	messageRepository := repository.NewMessageRepository()
	messageService := service.NewMessageService(messageRepository, db, validate)
	messageController := controller.NewMessageController(messageService)

	router := httprouter.New()

	router.GET("/api/messages", messageController.FindAll)
	router.GET("/api/messages/:messageId", messageController.FindById)
	router.POST("/api/messages", messageController.Create)
	router.DELETE("/api/messages/:messageId", messageController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
