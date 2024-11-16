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
	avg := 100
	if len(ans) > 0 {
		sum := 0
		for _, num := range ans {
			sum += num
		}
		avg = sum / len(ans)
	}
	desc := "! Отличная совместимость со всеми участниками команды." +
		" Личностные качества и профессиональные навыки кандидата идеально дополняют команду, создавая гармоничную и продуктивную рабочую атмосферу. "
	if avg <= 65 {
		desc = "Кандидат демонстрирует средний уровень совместимости с командой. " +
			"Его профессиональные навыки и личные качества имеют потенциал для успешного взаимодействия," +
			" но могут потребовать дополнительной адаптации или проработки в некоторых аспектах."
	}
	if avg <= 40 {
		desc = "К сожалению, кандидат демонстрирует низкий уровень совместимости с командой. Его профессиональные качества и стиль взаимодействия" +
			" существенно отличаются от текущей динамики коллектива, что может затруднить эффективное сотрудничество. "
	}
	c.JSON(http.StatusOK, compareLotRes{
		Percents:    ans,
		Avg:         avg,
		Description: desc,
	})
}

type compare2Req struct {
	Date1 time.Time `json:"date1" binding:"required"`
	Date2 time.Time `json:"date2" binding:"required"`
}

type compareLotReq struct {
	Participants []entities.Participant `json:"patrticipants"`
	BirthDate    time.Time              `json:"birth_date"`
}

type compareLotRes struct {
	Percents    []int  `json:"percents"`
	Avg         int    `json:"avg"`
	Description string `json:"description"`
}
