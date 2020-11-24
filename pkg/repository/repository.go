package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sashindionicus/rest-api"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
