package commands

import (
	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/application/ports"
	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/domain"
)

type MarkTodoDone struct {
	TodoID string `json:"todo_id"`
}
type MarkTodoDoneHandler struct {
	TodosR ports.TodoReader
	TodosW ports.TodoWriter
}

func (cmdHandler MarkTodoDoneHandler) Handle(cmd MarkTodoDone) error {
	if cmd.TodoID == "" {
		return domain.ValidationError{Msg: "todo_id is required"}
	}
	todo, _ := cmdHandler.TodosR.GetByID(cmd.TodoID)
	if todo == nil {
		return domain.TodoNotFound{TodoID: cmd.TodoID}
	}
	updated := todo.MarkDone()
	return cmdHandler.TodosW.Update(updated)
}
