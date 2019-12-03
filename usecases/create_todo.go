package usecases

import (
	"context"

	"github.com/micheam/clean-architecture-go/entities"
	"github.com/micheam/clean-architecture-go/usecases/interfaces"
	"github.com/pkg/errors"
)

// Interactor
type (
	CreateTodoInteractor struct {
		Registerer interfaces.TodoRegisterer
		OutputPort CreateTodoOutputPort
	}
)

// Input Port
type (
	CreateTodoInputPort interface {
		Handle(ctx context.Context, data CreateTodoInputData) error
	}
	CreateTodoInputData struct {
		Title       string
		Description *string
	}
)

// Output Port
type (
	CreateTodoOutputPort interface {
		Complete(ctx context.Context, data *CreateTodoOutputData) error
	}
	CreateTodoOutputData struct {
		Saved interfaces.SavedTodo
	}
)

func (c CreateTodoInteractor) Handle(ctx context.Context, input CreateTodoInputData) error {

	var err error

	var title entities.Title
	if title, err = entities.NewTitle(input.Title); err != nil {
		return NewErrIllegalInputData(err.Error())
	}
	var unsaved = interfaces.UnsavedTodo{
		ID: entities.NewID(), Title: title,
	}

	if input.Description != nil {
		var desc entities.Description
		if desc, err = entities.NewDescription(*input.Description); err != nil {
			return NewErrIllegalInputData(err.Error())
		}
		unsaved.Description = &desc
	}

	var saved interfaces.SavedTodo
	if saved, err = c.Registerer.Register(ctx, unsaved); err != nil {
		return errors.Wrap(err, "Failed to regiser todo")
	}

	var output = &CreateTodoOutputData{Saved: saved}
	return c.OutputPort.Complete(ctx, output)
}
