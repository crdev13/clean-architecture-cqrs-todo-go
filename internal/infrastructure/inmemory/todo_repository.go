package inmemory

import (
	"sync"

	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/domain"
)

type TodoRepository struct {
	mu   sync.RWMutex
	byID map[string]domain.Todo
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		byID: make(map[string]domain.Todo),
	}
}

func (repo *TodoRepository) GetByID(todoID string) (*domain.Todo, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	todo, ok := repo.byID[todoID]
	if !ok {
		return nil, nil
	}
	todoCopy := todo

	return &todoCopy, nil
}

func (repo *TodoRepository) ListByUser(userID string) ([]domain.Todo, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	todoList := make([]domain.Todo, 0)
	for _, todo := range repo.byID {
		if todo.UserID == userID {
			todoList = append(todoList, todo)
		}
	}

	return todoList, nil
}

func (repo *TodoRepository) Save(todo domain.Todo) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.byID[todo.ID] = todo
	return nil
}

func (repo *TodoRepository) Update(todo domain.Todo) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.byID[todo.ID] = todo
	return nil
}
