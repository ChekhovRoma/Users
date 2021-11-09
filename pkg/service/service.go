package service

import (
	"context"
	"users/models"
	"users/pkg/repository"
)

type Authorization interface {
	SignUp(name, email, password, role string) (int, error)
	ParseToken(token string) (int, error)
	SignIn(ctx context.Context, email, password string) (models.Tokens, error)
}

type Services struct {
	Authorization Authorization
	TokenManager  TokenManager
}

func NewService(repos *repository.Repositories) *Services {
	return &Services{
		Authorization: NewAuthorizationService(repos.UserRepo, NewManager(signatureKey)),
	}
}
