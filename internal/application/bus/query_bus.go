package bus

import "context"

type QueryBus struct {
	ListUserTodos func(context.Context, any) (any, error)
}
