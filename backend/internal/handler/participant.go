package handler

import (
	"github.com/fullstakilla/naimix-hackathon/backend/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// AddParticipant godoc
// @Summary      Add a participant
// @Description  Add a participant to the database.
// @Tags         participant
// @Accept       json
// @Produce      json
// @Param        input  body      participantReq  true  "Participant data"
// @Success      200    {object}  map[string]interface{}  "status"="ok"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/participant/add [post]
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
	Id         int       `json:"id"`
	TeamId     int       `json:"team_id"`
	Name       string    `json:"name" `
	Role       string    `json:"role"`
	BirthDate  time.Time `json:"birthdate"`
	BirthPlace string    `json:"birthplace"`
}

// GetAllParticipants godoc
// @Summary      Get all participants
// @Description  Get all participants from the database.
// @Tags         participant
// @Accept       json
// @Produce      json
// @Success      200    {object}  []entities.Participant  "participants"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/participant/all [get]
func (h *Handler) GetAllParticipants(c *gin.Context) {
	participants, err := h.service.GetAllParticipants(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, participants)
}

// GetParticipantById godoc
// @Summary      Get a participant by ID
// @Description  Get a participant from the database by its ID.
// @Tags         participant
// @Accept       json
// @Produce      json
// @Param        id     path      int  true  "Participant ID"
// @Success      200    {object}  entities.Participant  "participant"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/participant/{id} [get]
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

// GetParticipantsByTeamId godoc
// @Summary      Get participants by team ID
// @Description  Get all participants from the database by team ID.
// @Tags         participant
// @Accept       json
// @Produce      json
// @Param        id     path      int  true  "Team ID"
// @Success      200    {object}  []entities.Participant  "participants"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/participant/team/{id} [get]
func (h *Handler) GetParticipantsByTeamId(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	userId, exists := c.Get("userID")
	if !exists {
		newErrorResponse(c, http.StatusUnauthorized, "user not authenticated")
		return
	}

	isCreator, err := h.service.Team.IsUserTeamCreator(c, userId.(int), id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if !isCreator {
		newErrorResponse(c, http.StatusForbidden, "user is not the creator of the team")
		return
	}

	participants, err := h.service.GetParticipantsByTeamId(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, participants)
}

// UpdateParticipant godoc
// @Summary      Update a participant
// @Description  Update a participant in the database.
// @Tags         participant
// @Accept       json
// @Produce      json
// @Param        id     path      int            true  "Participant ID"
// @Param        input  body      participantReq  true  "Participant data"
// @Success      200    {object}  map[string]interface{}  "status"="ok"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/participant/{id} [put]
func (h *Handler) UpdateParticipant(c *gin.Context) {
	var input participantReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.service.UpdateParticipant(c, entities.Participant{Name: input.Name, TeamID: input.TeamId, Role: input.Role, BirthDate: input.BirthDate, BirthPlace: input.BirthPlace, Id: input.Id})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// DeleteParticipantById godoc
// @Summary      Delete a participant by ID
// @Description  Delete a participant from the database by its ID.
// @Tags         participant
// @Accept       json
// @Produce      json
// @Param        id     path      int  true  "Participant ID"
// @Success      200    {object}  map[string]interface{}  "status"="ok"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/participant/{id} [delete]
func (h *Handler) DeleteParticipantById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.DeleteParticipantById(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

// GetParticipantByCreatorId godoc
// @Summary      Get participants by creator ID
// @Description  Get all participants from the database by creator ID.
// @Tags         participant
// @Accept       json
// @Produce      json
// @Param        id     path      int  true  "Creator ID"
// @Success      200    {object}  []entities.Participant  "participants"
// @Failure      400    {object}  map[string]string "error"="Bad request"
// @Failure      500    {object}  map[string]string "error"="Internal server error"
// @Router       /api/v1/participant/creator/{id} [get]
func (h *Handler) GetParticipantByCreatorId(c *gin.Context) {
	creatorIdStr := c.Param("id")
	creatorId, err := strconv.Atoi(creatorIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid creator_id param")
		return
	}

	userId, exists := c.Get("userID")
	if !exists {
		newErrorResponse(c, http.StatusUnauthorized, "user not authenticated")
		return
	}

	if userId.(int) != creatorId {
		newErrorResponse(c, http.StatusForbidden, "user is not the creator")
		return
	}

	participants, err := h.service.GetParticipantByCreatorId(c, creatorId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, participants)
}
