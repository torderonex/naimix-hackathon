package handler

import (
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// signup godoc
// @Summary      Register a new user
// @Description  Create a new user account and generate a JWT token.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      loginDto  true  "User credentials"
// @Success      200    {object}  map[string]interface{}  "token"="JWT token"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/auth/signup [post]
func (h *Handler) register(c *gin.Context) {
	var input loginDto
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err := h.service.Authorization.CreateUser(c, entities.User{Username: input.Username, Password: input.Password})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	token, err := h.service.Authorization.GenerateToken(c, input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

type loginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// login godoc
// @Summary      Log in a user
// @Description  Authenticate a user and generate a JWT token.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      loginDto  true  "User credentials"
// @Success      200    {object}  map[string]interface{}  "token"="JWT token"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var input loginDto

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(c, input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
