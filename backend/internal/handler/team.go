package handler

import (
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddTeam godoc
// @Summary      Add a team
// @Description  Add a team to the database.
// @Tags         team
// @Accept       json
// @Produce      json
// @Param        input  body      teamReq  true  "Team data"
// @Success      200    {object}  map[string]interface{}  "status"="ok"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/team/add [post]
func (h *Handler) AddTeam(c *gin.Context) {
	var input teamReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err := h.service.CreateTeam(c, entities.Team{Name: input.Name, Description: input.Description, UserID: input.User})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

type teamReq struct {
	User        int    `json:"user_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Id          int    `json:"id"`
}

// GetAllTeams godoc
// @Summary      Get all teams
// @Description  Get all teams from the database.
// @Tags         team
// @Accept       json
// @Produce      json
// @Success      200    {object}  []entities.Team  "teams"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/team/all [get]
func (h *Handler) GetAllTeams(c *gin.Context) {
	teams, err := h.service.GetAllTeams(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, teams)
}

// GetTeamById godoc
// @Summary      Get a team by ID
// @Description  Get a team from the database by its ID.
// @Tags         team
// @Accept       json
// @Produce      json
// @Param        id     path      int  true  "Team ID"
// @Success      200    {object}  entities.Team  "team"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/team/{id} [get]
func (h *Handler) GetTeamById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	team, err := h.service.GetTeamById(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, team)
}

// GetTeamsByUserId godoc
// @Summary      Get teams by user ID
// @Description  Get all teams from the database by user ID.
// @Tags         team
// @Accept       json
// @Produce      json
// @Param        id     path      int  true  "User ID"
// @Success      200    {object}  []entities.Team  "teams"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/team/user/{id} [get]
func (h *Handler) GetTeamsByUserId(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	teams, err := h.service.GetTeamsByUserId(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, teams)
}

// UpdateTeam godoc
// @Summary      Update a team
// @Description  Update a team in the database.
// @Tags         team
// @Accept       json
// @Produce      json
// @Param        id     path      int  true  "Team ID"
// @Param        input  body      teamReq  true  "Team data"
// @Success      200    {object}  map[string]interface{}  "status"="ok"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/team/{id} [put]
func (h *Handler) UpdateTeam(c *gin.Context) {
	var input teamReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.UpdateTeam(c, entities.Team{Name: input.Name, Description: input.Description, Id: input.Id})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// DeleteTeamById godoc
// @Summary      Delete a team by ID
// @Description  Delete a team from the database by its ID.
// @Tags         team
// @Accept       json
// @Produce      json
// @Param        id     path      int  true  "Team ID"
// @Success      200    {object}  map[string]interface{}  "status"="ok"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/team/{id} [delete]
func (h *Handler) DeleteTeamById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.DeleteTeamById(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
