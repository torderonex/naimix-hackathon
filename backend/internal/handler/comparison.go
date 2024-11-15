package handler

import (
	"github.com/fullstakilla/naimix-hackathon/backend/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) compare(c *gin.Context) {
	var input compareReq
	if err := c.BindJSON(&input); err != nil {
		// Если ошибка валидации запроса
		newErrorResponse(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	// Вызов сервиса для сравнения
	ll, err := service.Test(input.Text)
	if err != nil {
		// Если сервис возвращает ошибку
		newErrorResponse(c, http.StatusInternalServerError, "Comparison failed: "+err.Error())
		return
	}

	// Возвращаем результат
	c.JSON(http.StatusOK, ll)
}

// compareReq структура запроса для сравнения текста
type compareReq struct {
	Text string `json:"text" binding:"required"` // Текст, который будет сравниваться
}
