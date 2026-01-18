package ports

import "github.com/crdev13/clean-architecture-cqrs-todo-go/internal/domain"

type TodoReader interface {
	ListByUser(userID string) ([]domain.Todo, error)
	GetById(todoID string) (*domain.Todo, error)
}
