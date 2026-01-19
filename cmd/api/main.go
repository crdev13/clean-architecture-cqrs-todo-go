package main

import (
	"context"
	"log"
	"net/http"

	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/application/bus"
	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/application/commands"
	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/application/queries"
	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/infrastructure/inmemory"
	ihttp "github.com/crdev13/clean-architecture-cqrs-todo-go/internal/interface/http"
)

func main() {
	// In-memory adapters
	userRepo := inmemory.NewUserRepository()
	todoRepo := inmemory.NewTodoRepository()

	// Command handlers
	createUserH := commands.CreateUserHandler{UsersR: userRepo, UsersW: userRepo}
	createTodoH := commands.CreateTodoHandler{UsersR: userRepo, TodosW: todoRepo}
	markDoneH := commands.MarkTodoDoneHandler{TodosR: todoRepo, TodosW: todoRepo}

	// Query handlers
	listTodosH := queries.ListUserTodosHandler{UsersR: userRepo, TodosR: todoRepo}

	// Buses
	cmdBus := bus.CommandBus{
		CreateUser: func(ctx context.Context, c any) (any, error) {
			return createUserH.Handle(c.(commands.CreateUser))
		},
		CreateTodo: func(ctx context.Context, c any) (any, error) {
			return createTodoH.Handle(c.(commands.CreateTodo))
		},
		MarkTodoDone: func(ctx context.Context, c any) (any, error) {
			return nil, markDoneH.Handle(c.(commands.MarkTodoDone))
		},
	}
	qryBus := bus.QueryBus{
		ListUserTodos: func(ctx context.Context, q any) (any, error) {
			return listTodosH.Handle(q.(queries.ListUserTodos))
		},
	}

	handlers := ihttp.Handlers{Cmd: cmdBus, Qry: qryBus}
	mux := ihttp.NewServer(handlers)

	log.Println("listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
