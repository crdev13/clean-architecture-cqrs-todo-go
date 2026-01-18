package domain

import "time"

type Todo struct {
	ID        string
	UserID    string
	Title     string
	Done      bool
	CreatedAt time.Time
}
