package ports

import "github.com/crdev13/clean-architecture-cqrs-todo-go/internal/domain"

type TodoWriter interface {
	Save(todo domain.Todo) error
	Update(todo domain.Todo) error
}
