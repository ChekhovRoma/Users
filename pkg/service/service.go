package service

import (
	"context"
	"users/pkg/handler"
	"users/pkg/repository"
)

//todo move to config
const salt = "edrftgyhujikiuy"

type Authorization interface {
	SignUp(name, email, password, role string) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
	SignIn(ctx context.Context, email, password string) (handler.Tokens, error)
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
