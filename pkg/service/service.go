package service

import (
	"users/pkg/repository"
)

type Authorization interface {
	Create(email, password string) (int, error)
}

type Services struct {
	Authorization Authorization
}

func NewService(repos *repository.Repositories) *Services {
	return &Services{
		Authorization: NewAuthorizationService(repos.UserRepo),
	}
}
