// Package inmemory that implements user repository
package inmemory

import (
	"sync"

	"github.com/crdev13/clean-architecture-cqrs-todo-go/internal/domain"
)

type UserRepository struct {
	mu      sync.RWMutex
	byID    map[string]domain.User
	byEmail map[string]string
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		byID:    make(map[string]domain.User),
		byEmail: make(map[string]string),
	}
}

func (repo *UserRepository) GetByID(userID string) (*domain.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	user, ok := repo.byID[userID]
	if !ok {
		return nil, nil
	}
	userCopy := user

	return &userCopy, nil
}

func (repo *UserRepository) GetByEmail(email string) (*domain.User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	userID, ok := repo.byEmail[email]
	if !ok {
		return nil, nil
	}
	user := repo.byID[userID]
	userCopy := user

	return &userCopy, nil
}

func (repo *UserRepository) Save(user domain.User) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.byID[user.ID] = user
	repo.byEmail[user.Email] = user.ID
	return nil
}
