package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"users/models"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		fmt.Println("400")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		fmt.Println("400 ", err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
