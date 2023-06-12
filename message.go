package gonotifier

import "github.com/gofrs/uuid"

type Status int

const (
	Success Status = iota + 1
	Warning
	Error
)

type Message struct {
	UserID uuid.UUID `json:"userId"`
	Status Status    `json:"status"`
	Header string    `json:"header"`
	Body   string    `json:"body"`
}
