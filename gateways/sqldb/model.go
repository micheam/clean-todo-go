package sqldb

import (
	"github.com/google/uuid"
	"github.com/micheam/clean-architecture-go/entities"
	"github.com/micheam/clean-architecture-go/usecases/interfaces"
	"github.com/pkg/errors"
	"time"
)

type Todo struct {
	ID          string  `db:"id"`
	Title       string  `db:"title"`
	Description *string `db:"description"`
	Done        bool    `db:"done"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
}

func (r Todo) ToSavedTodo() (saved interfaces.SavedTodo, err error) {

	id, err := uuid.Parse(r.ID)
	if err != nil {
		err = errors.Wrap(err, "failed to parse todo_id")
		return
	}

	createdAt, err := time.Parse(time.RFC3339, r.CreatedAt)
	if err != nil {
		err = errors.Wrap(err, "failed to parse created_at")
		return
	}

	updatedAt, err := time.Parse(time.RFC3339, r.UpdatedAt)
	if err != nil {
		err = errors.Wrap(err, "failed to parse updated_at")
		return
	}

	saved = interfaces.SavedTodo{
		UnsavedTodo: interfaces.UnsavedTodo{
			ID:    entities.ID(id),
			Title: entities.Title(r.Title),
		},
		Done:      r.Done,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	if r.Description != nil {
		desc := entities.Description(*r.Description)
		saved.Description = &desc
	}

	return
}

type TodoCollection []Todo

func (t TodoCollection) ToSavedTodo() ([]interfaces.SavedTodo, error) {
	var result = make([]interfaces.SavedTodo, 0)
	for i, r := range ([]Todo)(t) {
		e, err := r.ToSavedTodo()
		if err != nil {
			return result, errors.Wrapf(err, "failed to convert row(%v) at %d", r, i)
		}
		result = append(result, e)
	}
	return result, nil
}
