package interfaces

import (
	"github.com/micheam/clean-architecture-go/entities"
	"time"
)

type UnsavedTodo struct {
	ID          entities.ID
	Title       entities.Title
	Description *entities.Description
}

type SavedTodo struct {
	UnsavedTodo
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
