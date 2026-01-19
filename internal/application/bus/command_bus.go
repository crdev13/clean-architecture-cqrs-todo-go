// Package bus that implements cqrs
package bus

import "context"

type CommandBus struct {
	CreateUser   func(context.Context, any) (any, error)
	CreateTodo   func(context.Context, any) (any, error)
	MarkTodoDone func(context.Context, any) (any, error)
}
