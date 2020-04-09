package registry

import (
	"github.com/y-shika/todo_golang_api/domain/repository"
	"github.com/y-shika/todo_golang_api/stub"
	"github.com/y-shika/todo_golang_api/usecase"
)

// NewTodoUseCase creates new usecase.Todo.
func NewTodoUseCase() *usecase.Todo {
	return usecase.NewTodo(newTodoRepository())
}

func newTodoRepository() repository.Todo {
	// TODO: 一旦動作確認のためにstubから取得するようにする。確認が出来たら元に戻す (configも)。
	// if config.IsMockMode() {
	// 	return stub.NewTodo()
	// }
	// return gateway.NewTodo()
	return stub.NewTodo()
}
