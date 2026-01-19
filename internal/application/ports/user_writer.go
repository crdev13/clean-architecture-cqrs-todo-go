package ports

import "github.com/crdev13/clean-architecture-cqrs-todo-go/internal/domain"

type UserWriter interface {
	Save(user domain.User) error
}
