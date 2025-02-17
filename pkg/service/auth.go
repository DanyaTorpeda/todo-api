package service

import (
	todolist "todo-app"
	"todo-app/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Create(input todolist.User) (int, error) {
	hashedPassword, err := generatePassword(input.Password)
	if err != nil {
		return 0, err
	}
	input.Password = hashedPassword
	return s.repo.Create(input)
}

func (s *AuthService) GenerateSession(input todolist.User) (int, error) {
	user, err := s.repo.Get(input.Username)
	if err != nil {
		return 0, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return 0, err
	}

}

func generatePassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return (string(hashed)), nil
}
