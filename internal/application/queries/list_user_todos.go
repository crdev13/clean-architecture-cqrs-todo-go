// Package queries that implements cqrs
package queries

import (
	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/application/ports"
	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/domain"
)

type ListUserTodos struct {
	UserID string
}

type ListUserTodosHandler struct {
	UsersR ports.UserReader
	TodosR ports.TodoReader
}

func (queryHandler ListUserTodosHandler) Handle(query ListUserTodos) ([]domain.Todo, error) {
	user, _ := queryHandler.UsersR.GetByID(query.UserID)
	if user == nil {
		return nil, domain.UserNotFound{UserID: query.UserID}
	}
	return queryHandler.TodosR.ListByUser(query.UserID)
}
