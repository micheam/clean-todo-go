package sqldb

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/micheam/clean-architecture-go/entities"
	"github.com/micheam/clean-architecture-go/usecases/interfaces"
	"github.com/pkg/errors"
)

var _ interfaces.TodoRepository = SQLDBTodoRepository{}
var _ interfaces.TodoRegisterer = SQLDBTodoRepository{}
var _ interfaces.TodoGetter = SQLDBTodoRepository{}
var _ interfaces.TodoFinder = SQLDBTodoRepository{}

type SQLDBTodoRepository struct {
	DB *sqlx.DB
}

func (s SQLDBTodoRepository) Register(ctx context.Context, unsaved interfaces.UnsavedTodo) (
	saved interfaces.SavedTodo, err error) {

	var desc *string
	if unsaved.Description != nil {
		s := unsaved.Description.String()
		desc = &s
	}

	var (
		now      time.Time = time.Now()
		nowFmted string    = now.Format(time.RFC3339)
	)

	row := Todo{
		ID:          unsaved.ID.String(),
		Title:       unsaved.Title.String(),
		Description: desc,
		Done:        false,
		CreatedAt:   nowFmted,
		UpdatedAt:   nowFmted,
	}

	if _, err = s.DB.NamedExecContext(ctx,
		`INSERT INTO todo (id, title, description, done, created_at, updated_at)
		VALUES (:id, :title, :description, :done, :created_at, :updated_at);`, row); err != nil {
		err = errors.Wrap(err, "failed to Register new todo")
		return
	}

	saved.UnsavedTodo = unsaved
	saved.Done = false
	saved.CreatedAt = now
	saved.UpdatedAt = now
	return
}

const selectTodoClause = `
SELECT 
  id, 
  title, 
  description, 
  done, 
  created_at, 
  updated_at 
FROM 
  todo `

func (s SQLDBTodoRepository) Get(ctx context.Context, id entities.ID) (saved interfaces.SavedTodo, found bool, err error) {

	rows := make([]Todo, 0)

	if err = s.DB.SelectContext(ctx, &rows,
		selectTodoClause+`WHERE id=?;`, id.String()); err != nil {
		err = errors.Wrap(err, "failed to select row from todo")
		return
	}

	if len(rows) == 0 {
		found = false
		return
	}

	saved, err = rows[0].ToSavedTodo()
	if err != nil {
		err = errors.Wrap(err, "failed to convert row to entity")
		return
	}

	found = true
	return
}

func (s SQLDBTodoRepository) List(ctx context.Context) ([]interfaces.SavedTodo, error) {

	rows := make([]Todo, 0)
	if err := s.DB.SelectContext(ctx, &rows,
		selectTodoClause); err != nil {
		return nil, errors.Wrap(err, "failed to select row from todo")
	}

	list, err := (TodoCollection)(rows).ToSavedTodo()
	if err != nil {
		return nil, err
	}

	return list, nil
}
