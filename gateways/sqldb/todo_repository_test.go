package sqldb

import (
	"context"
	"testing"

	"github.com/micheam/clean-architecture-go/entities"
	"github.com/micheam/clean-architecture-go/usecases/interfaces"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestSQLDBTodoRepository_Register(t *testing.T) {
	sut := SQLDBTodoRepository{DB: testdb}
	ctx := context.TODO()

	var (
		title   entities.Title         = entities.Title("THIS IS A TITLE")
		unsaved interfaces.UnsavedTodo = interfaces.UnsavedTodo{
			ID: entities.NewID(), Title: title,
		}
	)

	got, err := sut.Register(ctx, unsaved)
	if assert.NoError(t, err) {
		assert.Equal(t, title, got.Title)
		assert.Nil(t, got.Description)
		assert.NotNil(t, got.CreatedAt)
		assert.NotNil(t, got.UpdatedAt)
	}
}

func TestSQLDBTodoRepository_Register_With_OptionalVal(t *testing.T) {
	sut := SQLDBTodoRepository{DB: testdb}
	ctx := context.TODO()

	var (
		title   entities.Title         = entities.Title("THIS IS A TITLE")
		desc    entities.Description   = entities.Description("THIS IS A DESCRIPTION")
		unsaved interfaces.UnsavedTodo = interfaces.UnsavedTodo{
			ID: entities.NewID(), Title: title,
		}
	)
	unsaved.Description = &desc

	got, err := sut.Register(ctx, unsaved)
	if assert.NoError(t, err) {
		assert.Equal(t, title, got.Title)
		assert.Equal(t, &desc, got.Description)

		// timestamps will set inside method
		assert.NotNil(t, got.CreatedAt)
		assert.NotNil(t, got.UpdatedAt)
	}
}

func TestSQLDBTodoRepository_Get(t *testing.T) {
	sut := SQLDBTodoRepository{DB: testdb}
	ctx := context.TODO()

	var (
		title   entities.Title         = entities.Title("THIS IS A TITLE")
		unsaved interfaces.UnsavedTodo = interfaces.UnsavedTodo{
			ID: entities.NewID(), Title: title,
		}
	)

	saved, err := sut.Register(ctx, unsaved)
	assert.NoError(t, err)

	got, found, err := sut.Get(ctx, saved.ID)
	if assert.NoError(t, err) {
		assert.True(t, found)
		assert.Equal(t, saved.ID, got.ID)
		assert.Equal(t, saved.Title, got.Title)
		assert.Equal(t, saved.Description, got.Description)
		assert.Equal(t, saved.Done, got.Done)

		assert.Equal(t, saved.CreatedAt.Format(time.RFC3339), got.CreatedAt.Format(time.RFC3339))
		assert.Equal(t, saved.UpdatedAt.Format(time.RFC3339), got.UpdatedAt.Format(time.RFC3339))
	}
}

func TestSQLDBTodoRepository_Get_NoExist_Todo(t *testing.T) {
	sut := SQLDBTodoRepository{DB: testdb}
	ctx := context.TODO()

	_, found, err := sut.Get(ctx, entities.NewID())

	if assert.NoError(t, err) {
		assert.False(t, found)
	}
}

func TestSQLDBTodoRepository_List(t *testing.T) {
	sut := SQLDBTodoRepository{DB: testdb}
	ctx := context.TODO()

	var (
		title1 entities.Title = entities.Title("THIS IS A TITLE 1")
		title2 entities.Title = entities.Title("THIS IS A TITLE 2")
	)

	_, _ = sut.Register(ctx, interfaces.UnsavedTodo{ID: entities.NewID(), Title: title1})
	_, _ = sut.Register(ctx, interfaces.UnsavedTodo{ID: entities.NewID(), Title: title2})

	got, err := sut.List(ctx)
	if assert.NoError(t, err) {
		assert.LessOrEqual(t, 2, len(got))
	}
}
