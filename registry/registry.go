package registry

import (
	"github.com/y-shika/todo_golang_api/config"
	"github.com/y-shika/todo_golang_api/domain/repository"
	"github.com/y-shika/todo_golang_api/gateway"
	"github.com/y-shika/todo_golang_api/stub"
	"github.com/y-shika/todo_golang_api/usecase"
)

// NewTodoUseCase creates new usecase.Todo.
func NewTodoUseCase() *usecase.Todo {
	return usecase.NewTodo(newTodoRepository())
}

func newTodoRepository() repository.Todo {
	if config.IsMockMode() {
		return stub.NewTodo()
	}
	return gateway.NewTodo()
}
