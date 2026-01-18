// Package ports for application ports
package ports

import "github.com/crdev13/clean-architecture-cqrs-todo-go/internal/domain"

type UserReader interface {
	GetByID(userID string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
}
