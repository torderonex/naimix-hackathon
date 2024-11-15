package handler

import (
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) AddParticipant(c *gin.Context) {
	var input participantReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	_, err := h.service.CreateParticipant(c, entities.Participant{Name: input.Name, TeamID: input.TeamId, Role: input.Role, BirthDate: input.BirthDate, BirthPlace: input.BirthPlace})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

type participantReq struct {
	TeamId     int    `json:"team_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Role       string `json:"role" binding:"required"`
	BirthDate  string `json:"birthdate" binding:"required"`
	BirthPlace string `json:"birthplace" binding:"required"`
}

func (h *Handler) GetAllParticipants(c *gin.Context) {
	participants, err := h.service.GetAllParticipants(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, participants)
}

func (h *Handler) GetParticipantById(c *gin.Context) {
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

	participant, err := h.service.GetParticipantById(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, participant)
}

func (h *Handler) GetParticipantsByTeamId(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	participants, err := h.service.GetParticipantsByTeamId(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, participants)
}

func (h *Handler) UpdateParticipant(c *gin.Context) {
	var input participantReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.service.UpdateParticipant(c, entities.Participant{Name: input.Name, TeamID: input.TeamId, Role: input.Role, BirthDate: input.BirthDate, BirthPlace: input.BirthPlace})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
