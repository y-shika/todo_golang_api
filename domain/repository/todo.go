package repository

import (
	"context"

	"github.com/y-shika/todo_golang_api/domain/model"
)

// Todo is repository interface.
type Todo interface {
	// TODO: gatewayで実装するメソッドのinterfaceを定義する
	//       名前は仮決め
	GetTodoLists(ctx context.Context) ([]*model.Todo, error)
}
