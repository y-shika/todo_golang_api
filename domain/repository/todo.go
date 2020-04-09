package repository

import (
	"context"

	"github.com/y-shika/todo_golang_api/domain/model"
)

// Todo is repository interface.
type Todo interface {
	GetTodoLists(ctx context.Context) ([]*model.Todo, error)
}
