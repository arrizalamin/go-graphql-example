package todo

import (
	"github.com/google/uuid"
)

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func New(text string) Todo {
	return Todo{
		ID:   uuid.New().String(),
		Text: text,
		Done: false,
	}
}
