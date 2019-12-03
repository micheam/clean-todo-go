package usecases

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/micheam/clean-architecture-go/usecases/interfaces"
	"github.com/micheam/clean-architecture-go/utils/matcher"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodoInteractor_Handle(t *testing.T) {
	ctx := context.TODO()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		title = "This is a valid title"
		desc  = `# header`
		input = CreateTodoInputData{
			Title:       title,
			Description: &desc,
		}
		saved = interfaces.SavedTodo{}
	)

	registerer := interfaces.NewMockTodoRegisterer(ctrl)
	registerer.EXPECT().
		Register(ctx, matcher.OfType("interfaces.UnsavedTodo")).
		Return(saved, nil).
		Times(1)

	presenter := NewMockCreateTodoOutputPort(ctrl)
	presenter.EXPECT().
		Complete(ctx, gomock.Any()).
		Return(nil).
		Times(1)

	sut := CreateTodoInteractor{
		Registerer: registerer,
		OutputPort: presenter,
	}

	err := sut.Handle(ctx, input)
	assert.NoError(t, err)
}

func TestCreateTodoInteractor_Handle_Invalid_Title(t *testing.T) {
	ctx := context.TODO()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		title = "" // Illegal Title
		desc  = `# header`
		input = CreateTodoInputData{
			Title:       title,
			Description: &desc,
		}
	)

	// No Expection
	registerer := interfaces.NewMockTodoRegisterer(ctrl)
	presenter := NewMockCreateTodoOutputPort(ctrl)

	sut := CreateTodoInteractor{
		Registerer: registerer,
		OutputPort: presenter,
	}

	err := sut.Handle(ctx, input)
	if assert.Error(t, err) {
		assert.IsType(t, ErrIllegalInputData{}, err)
	}
}

func TestCreateTodoInteractor_Handle_Invalid_Description(t *testing.T) {
	ctx := context.TODO()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		title = "valid title"
		desc  = "-- THIS IS A VERY LONG DESCRIPTION ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------"
		input = CreateTodoInputData{
			Title:       title,
			Description: &desc,
		}
	)

	// No Expection
	registerer := interfaces.NewMockTodoRegisterer(ctrl)
	presenter := NewMockCreateTodoOutputPort(ctrl)

	sut := CreateTodoInteractor{
		Registerer: registerer,
		OutputPort: presenter,
	}

	err := sut.Handle(ctx, input)
	if assert.Error(t, err) {
		assert.IsType(t, ErrIllegalInputData{}, err)
	}
}
