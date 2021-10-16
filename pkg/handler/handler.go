package handler

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type AuthorizationService interface {
	SignUp(name, email, password, role string) (int, error)
	SignIn(ctx context.Context, email, password string) (Tokens, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Handler struct {
	as AuthorizationService
}

func NewHandler(as AuthorizationService) *Handler {
	return &Handler{as: as}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.POST("test", h.test)
	}

	return router
}
