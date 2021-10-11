package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input SignUpRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//todo validate data
	id, err := h.as.Create(input.Name, input.Email, input.Password, input.Role)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input SignInRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.as.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) test(c *gin.Context) {
	fmt.Println("ITS ALIVE!!!!!!!!")
}
