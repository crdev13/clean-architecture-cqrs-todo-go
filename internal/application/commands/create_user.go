// Package commands for the CQRS implementation
package commands

import (
	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/application/ports"
	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/domain"
	"github.com/google/uuid"
)

type CreateUser struct {
	Email string `json:"email"`
}

type CreateUserHandler struct {
	UsersR ports.UserReader
	UsersW ports.UserWriter
}

func (cmdHandler CreateUserHandler) Handle(cmd CreateUser) (string, error) {
	if cmd.Email == "" {
		return "", domain.ValidationError{Msg: "email is required"}
	}
	if existing, _ := cmdHandler.UsersR.GetByEmail(cmd.Email); existing != nil {
		return "", domain.EmailAlreadyUsed{Email: cmd.Email}
	}
	id := uuid.NewString()
	err := cmdHandler.UsersW.Save(domain.User{ID: id, Email: cmd.Email})
	if err != nil {
		return "", err
	}
	return id, nil
}
