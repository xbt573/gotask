package types

import (
	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Priority    int       `json:"priority,omitempty"`
}

type ListResponse struct {
	Ok      bool
	Message string
	Tasks   []Task
}

type DeleteRequest struct {
	Id uuid.UUID
}

type DeleteResponse struct {
	Ok      bool
	Message string
	Task
}

type InfoRequest struct {
	Id uuid.UUID
}

type InfoResponse struct {
	Ok      bool
	Message string
	Task
}

type NewRequest struct {
	Name        string
	Description string
	Priority    int
}

type NewResponse struct {
	Ok      bool
	Message string
	Task
}
