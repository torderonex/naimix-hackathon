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
// @Param        input  body      signupReq  true  "User credentials"
// @Success      200    {object}  map[string]interface{}  "token"="JWT token"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/auth/signup [post]
func (h *Handler) register(c *gin.Context) {
	var input signupReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err := h.service.Authorization.CreateUser(c, entities.User{Username: input.Username, Password: input.Password, Email: input.Email})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	token, err := h.service.Authorization.GenerateToken(c, input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

type signupReq struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginReq struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// login godoc
// @Summary      Log in a user
// @Description  Authenticate a user and generate a JWT token.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      loginReq  true  "User credentials"
// @Success      200    {object}  map[string]interface{}  "token"="JWT token"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var input loginReq

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(c, input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// getUserInfo godoc
// @Summary      Get user information
// @Description  Retrieve detailed information about the authenticated user.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200    {object}  map[string]interface{}  "User details"
// @Failure      401    {object}  map[string]string "error"="Unauthorized"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/auth/identity [get]
// @Security     ApiKeyAuth
func (h *Handler) getUserInfo(c *gin.Context) {
	userId, ok := c.Get("userID")
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user is unauthorized")
		return
	}
	info, err := h.service.GetUserById(c, userId.(int))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	info.Password = ""
	c.JSON(http.StatusOK, info)
}
