package domain

import "time"

type Todo struct {
	ID        string
	UserID    string
	Title     string
	Done      bool
	CreatedAt time.Time
}

func (t Todo) MarkDone() Todo {
	t.Done = true

	return t
}
