package repository

import (
	"fmt"
	todolist "todo-app"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) Create(input todolist.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES($1, $2, $3) RETURNING id",
		usersTable)

	var id int
	row := r.db.QueryRow(query, input.Name, input.Username, input.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) Get(username string) (todolist.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = $1", usersTable)
	var user todolist.User
	err := r.db.Get(&user, query, username)
	if err != nil {
		return todolist.User{}, err
	}

	return user, nil
}
