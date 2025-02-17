package service

import (
	todolist "todo-app"
	"todo-app/pkg/repository"
)

type Authorization interface {
	Create(input todolist.User) (int, error)
	GenerateSession(input todolist.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
