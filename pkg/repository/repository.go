package repository

import (
	todolist "todo-app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	Create(input todolist.User) (int, error)
	Get(username string) (todolist.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
