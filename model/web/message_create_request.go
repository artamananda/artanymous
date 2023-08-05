package web

type MessageCreateRequest struct {
	Question string `validate:"required"`
}
