package service

import (
	"users/pkg/repository"
)

type Authorization interface {
	Create(name, email, password, role string) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Services struct {
	Authorization Authorization
}

func NewService(repos *repository.Repositories) *Services {
	return &Services{
		Authorization: NewAuthorizationService(repos.UserRepo),
	}
}
