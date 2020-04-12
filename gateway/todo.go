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
	db, err := NewMySQLClient()
	if err != nil {
		return nil, fmt.Errorf("initialize MySQL client: %w", err)
	}
	defer db.Close()

	log.Println("Succeed to connect MySQL")

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		return nil, fmt.Errorf("get todos from MySQL: %w", err)
	}
	defer rows.Close()

	todos := make([]*model.Todo, 0)

	for rows.Next() {
		var todo model.Todo

		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Active, &todo.Detail); err != nil {
			return nil, fmt.Errorf("get record from todo table: %w", err)
		}

		todos = append(todos, &todo)
	}

	return todos, nil
}
