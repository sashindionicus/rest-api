package service

import (
	"github.com/sashindionicus/rest-api"
	"github.com/sashindionicus/rest-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
