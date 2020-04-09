package usecase

import (
	"context"
	"fmt"

	"github.com/y-shika/todo_golang_api/domain/model"
	"github.com/y-shika/todo_golang_api/domain/repository"
)

// Todo is an usecase for Todo.
type Todo struct {
	repo repository.Todo
}

// NewTodo creates new Todo.
func NewTodo(repo repository.Todo) *Todo {
	return &Todo{repo: repo}
}

// GetTodos returns list of todo.
func (uc *Todo) GetTodos(ctx context.Context) ([]*model.Todo, error) {
	// TODO: 実装する
	todos, err := uc.repo.GetTodoLists(ctx)
	if err != nil {
		return nil, fmt.Errorf("get todos: %w", err)
	}

	return todos, nil
}
