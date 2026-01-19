package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/application/bus"
	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/application/commands"
	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/application/queries"
	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/domain"
)

type Handlers struct {
	Cmd bus.CommandBus
	Qry bus.QueryBus
}

func (h Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body commands.CreateUser
	if err := readJSON(r, &body); err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid json"})
	}

	res, err := h.Cmd.CreateUser(context.Background(), body)
	if err != nil {
		writeDomainErr(w, err)
		return
	}

	writeJSON(w, 201, map[string]string{"id": res.(string)})
}

func (h Handlers) CreateTodo(w http.ResponseWriter, r *http.Request) {
	// POST /users/{id}/todos
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 3 || parts[0] != "users" || parts[2] != "todos" {
		writeJSON(w, 404, map[string]string{"error": "not found"})
		return
	}
	userID := parts[1]

	var body struct {
		Title string `json:"title"`
	}
	if err := readJSON(r, &body); err != nil {
		writeJSON(w, 400, map[string]string{"error": "invalid json"})
		return
	}

	cmd := commands.CreateTodo{UserID: userID, Title: body.Title}
	res, err := h.Cmd.CreateTodo(context.Background(), cmd)
	if err != nil {
		writeDomainErr(w, err)
		return
	}
	writeJSON(w, 201, map[string]any{"id": res.(string)})
}

func (h Handlers) ListUserTodos(w http.ResponseWriter, r *http.Request) {
	// GET /users/{id}/todos
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 3 || parts[0] != "users" || parts[2] != "todos" {
		writeJSON(w, 404, map[string]string{"error": "not found"})
		return
	}
	userID := parts[1]

	q := queries.ListUserTodos{UserID: userID}
	res, err := h.Qry.ListUserTodos(context.Background(), q)
	if err != nil {
		writeDomainErr(w, err)
		return
	}
	writeJSON(w, 200, res)
}

func (h Handlers) MarkTodoDone(w http.ResponseWriter, r *http.Request) {
	// POST /todos/{id}/done
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 3 || parts[0] != "todos" || parts[2] != "done" {
		writeJSON(w, 404, map[string]string{"error": "not found"})
		return
	}
	todoID := parts[1]

	_, err := h.Cmd.MarkTodoDone(context.Background(), commands.MarkTodoDone{TodoID: todoID})
	if err != nil {
		writeDomainErr(w, err)
		return
	}
	w.WriteHeader(204)
}

func writeDomainErr(w http.ResponseWriter, err error) {
	switch err.(type) {
	case domain.ValidationError:
		writeJSON(w, 400, map[string]string{"error": err.Error()})
	case domain.EmailAlreadyUsed:
		writeJSON(w, 409, map[string]string{"error": err.Error()})
	case domain.UserNotFound, domain.TodoNotFound:
		writeJSON(w, 404, map[string]string{"error": err.Error()})
	default:
		writeJSON(w, 500, map[string]string{"error": "internal error"})
	}
}
