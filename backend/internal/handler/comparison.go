package handler

import (
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// @Summary      Compare Two Cards
// @Description  Сравнивает две карты Таро и возвращает описание совместимости.
// @Tags         compare
// @Accept       json
// @Produce      json
// @Param        input  body  compare2Req  true  "Dates for two participants"
// @Success      200    {object}  string   "Description of compatibility"
// @Failure      400    {object}  map[string]string  "Invalid input"
// @Router       /compare/two [post]
func (h *Handler) compareTwoWithDescription(c *gin.Context) {
	var input compare2Req
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}
	ans, _ := h.service.Comparison.GetComparisonBetween2WithDescription(input.Date1, input.Date2)
	c.JSON(http.StatusOK, ans)
}

// @Summary      Compare Team Compatibility
// @Description  Сравнивает карту участника с несколькими другими и возвращает список совместимостей.
// @Tags         compare
// @Accept       json
// @Produce      json
// @Param        input  body  compareLotReq  true  "Participants and main date"
// @Success      200    {array}  entities.ComparisonResult  "List of compatibilities"
// @Failure      400    {object}  map[string]string  "Invalid input"
// @Router       /compare/lot [post]
func (h *Handler) compareLot(c *gin.Context) {
	var input compareLotReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}
	ans, _ := h.service.Comparison.GetTeamComparison(input.Participants, input.BirthDate)
	c.JSON(http.StatusOK, ans)
}

type compare2Req struct {
	Date1 time.Time `json:"date1" binding:"required"`
	Date2 time.Time `json:"date2" binding:"required"`
}

type compareLotReq struct {
	Participants []entities.Participant `json:"patrticipants"`
	BirthDate    time.Time              `json:"birth_date"`
}
