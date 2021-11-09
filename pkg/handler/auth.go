package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input SignUpRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "required field is missing")
		return
	}

	//todo validate data
	//if input.Email == "" {
	//	newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	//	return
	//}
	fmt.Println(input.Name)

	//todo layer dto
	id, err := h.as.SignUp(input.Name, input.Email, input.Password, input.Role)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var input SignInRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "required field is missing")
		return
	}

	res, err := h.as.SignIn(c.Request.Context(), input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  res.AccessToken,
		"refreshToken": res.AccessToken,
	})
}

func (h *Handler) test(c *gin.Context) {
	fmt.Println("ITS ALIVE!!!!!!!!")

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}
