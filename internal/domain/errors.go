// Package domain for custom error types that implement the Error interface
package domain

import "fmt"

type EmailAlreadyUsed struct{ Email string }

func (e EmailAlreadyUsed) Error() string { return fmt.Sprintf("email already used: %v", e.Email) }

type UserNotFound struct{ UserID string }

func (e UserNotFound) Error() string { return fmt.Sprintf("user not found: %v", e.UserID) }

type TodoNotFound struct{ TodoID string }

func (e TodoNotFound) Error() string { return fmt.Sprintf("todo not fond: %v", e.TodoID) }

type ValidationError struct{ Msg string }

func (e ValidationError) Error() string { return fmt.Sprintf("validation error: %s", e.Msg) }
