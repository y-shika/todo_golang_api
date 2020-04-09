package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/y-shika/todo_golang_api/registry"
	"github.com/y-shika/todo_golang_api/web/handler"
)

// New initializes a new http router.
func New() http.Handler {
	r := mux.NewRouter()
	// マッチするパスがないときのハンドラ
	r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

	todoUseCase := registry.NewTodoUseCase()
	todoHandler := handler.NewTodoServiceHandler(todoUseCase)

	todoService := r.PathPrefix("/TodoService").Subrouter()
	todoService.Handle("/GetTodos", todoHandler.GetTodos()).Methods(http.MethodGet)

	return r
}
