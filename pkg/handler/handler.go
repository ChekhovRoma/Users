package handler

import (
	"github.com/gin-gonic/gin"
)

type AuthorizationService interface {
	Create(name, email, password, role string) (int, error)
	GenerateToken(email, password string) (string, error)
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

	return router
}
