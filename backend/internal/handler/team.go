package handler

import (
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
}

func (h *Handler) GetAllTeams(c *gin.Context) {
	teams, err := h.service.GetAllTeams(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, teams)
}

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

func (h *Handler) UpdateTeam(c *gin.Context) {
	var input teamReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.UpdateTeam(c, entities.Team{Name: input.Name, Description: input.Description})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
