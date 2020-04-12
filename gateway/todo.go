package gateway

import (
	"context"
	"fmt"
	"log"

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

// GetTodoLists returns
func (gw *Todo) GetTodoLists(ctx context.Context) ([]*model.Todo, error) {
	// TODO: 実装する
	db, err := NewMySQLClient()
	if err != nil {
		return nil, fmt.Errorf("initialize MySQL client: %w", err)
	}
	defer db.Close()

	log.Println("Succeed to connect MySQL")

	return []*model.Todo{}, nil
}
