package commands

import (
	"time"

	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/application/ports"
	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/domain"
	"github.com/google/uuid"
)

type CreateTodo struct {
	UserID string `json:"user_id"`
	Title  string `json:"title"`
}
type CreateTodoHandler struct {
	UsersR ports.UserReader
	TodosW ports.TodoWriter
}

func (cmdHandler CreateTodoHandler) Handle(cmd CreateTodo) (string, error) {
	if cmd.UserID == "" {
		return "", domain.ValidationError{Msg: "user_id is required"}
	}
	if cmd.Title == "" {
		return "", domain.ValidationError{Msg: "title is required"}
	}

	user, _ := cmdHandler.UsersR.GetByID(cmd.UserID)
	if user == nil {
		return "", domain.UserNotFound{UserID: cmd.UserID}
	}

	id := uuid.NewString()
	todo := domain.Todo{
		ID:        id,
		UserID:    cmd.UserID,
		Title:     cmd.Title,
		Done:      false,
		CreatedAt: time.Now().UTC(),
	}
	if err := cmdHandler.TodosW.Save(todo); err != nil {
		return "", err
	}
	return id, nil
}
