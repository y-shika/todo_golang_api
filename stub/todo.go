package stub

import (
	"context"

	"github.com/y-shika/todo_golang_api/domain/model"
	"github.com/y-shika/todo_golang_api/domain/repository"
)

var _ repository.Todo = &Todo{}

// Todo implements repository.Todo interface.
type Todo struct{}

// NewTodo creates a new Todo gateway.
func NewTodo() *Todo {
	return &Todo{}
}

// GetTodoLists returns todos fixtures.
func (gw *Todo) GetTodoLists(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{
		{
			ID:     "1",
			Title:  "Title_1",
			Active: true,
			Detail: "Detail_1",
		},
		{
			ID:     "2",
			Title:  "Title_2",
			Active: true,
			Detail: "Detaile_2",
		},
	}, nil
}
