package entities

import (
	"time"
)

// Todo is a Entity
type Todo struct {
	ID          ID
	Title       Title
	Description *Description
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTodo(title Title) *Todo {
	return &Todo{
		ID:        NewID(),
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (t *Todo) SetDesc(desc Description) {
	t.Description = &desc
}
