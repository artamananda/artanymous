package web

type MessageCreateRequest struct {
	Question string `validate:"required" json:"question"`
}
