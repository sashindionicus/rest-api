package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/sashindionicus/rest-api"
	"github.com/sashindionicus/rest-api/pkg/repository"
)

const salt = "f48f5n742537o9f5n734f50"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user rest.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
