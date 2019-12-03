package interfaces

import (
	"context"
	"github.com/micheam/clean-architecture-go/entities"
)

type TodoRepository interface {
	TodoRegisterer
	TodoGetter
	TodoFinder
}

type TodoRegisterer interface {
	Register(ctx context.Context, todo UnsavedTodo) (SavedTodo, error)
}
type TodoGetter interface {
	Get(ctx context.Context, id entities.ID) (SavedTodo, bool, error)
}
type TodoFinder interface {
	List(ctx context.Context) ([]SavedTodo, error)
}
