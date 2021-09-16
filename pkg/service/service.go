package service

import (
	"users/models"
	"users/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
