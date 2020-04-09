package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/y-shika/todo_golang_api/usecase"
)

// TodoServiceHandler is a handler for Todo.
type TodoServiceHandler struct {
	usecase *usecase.Todo
}

// NewTodoServiceHandler returns a new TodoServiceHandler
func NewTodoServiceHandler(usecase *usecase.Todo) *TodoServiceHandler {
	return &TodoServiceHandler{
		usecase: usecase,
	}
}

// GetTodos is API of getting todos.
func (h *TodoServiceHandler) GetTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		todos, err := h.usecase.GetTodos(ctx)
		if err != nil {
			// TODO: logを自分で実装するか検討する
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}

		res, err := json.Marshal(todos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(res))
	}
}
