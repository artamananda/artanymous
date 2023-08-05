package web

type MessageCreateRequest struct {
	Question string `validate:"required,min=1" json:"question"`
}
