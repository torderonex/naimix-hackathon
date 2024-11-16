package handler

import (
	docs "github.com/fullstakilla/naimix-hackathon/backend/docs"
	"github.com/fullstakilla/naimix-hackathon/backend/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.Use(CORSMiddleware())

	v1 := router.Group("/api/v1/")

	auth := v1.Group("/auth")
	{
		auth.POST("/signup", h.register)
		auth.POST("/login", h.login)
		auth.GET("/identity", h.userIdentity, h.getUserInfo)
	}

	team := v1.Group("/team")
	{
		team.POST("/add", h.AddTeam)
		team.GET("/all", h.GetAllTeams)
		team.GET("/:id", h.GetTeamById)
		team.PUT("/:id", h.UpdateTeam)
		team.DELETE("/:id", h.DeleteTeamById)
		team.GET("/user/:id", h.GetTeamsByUserId)
	}

	participant := v1.Group("/participant")
	participant.Use(h.userIdentity)
	{
		participant.POST("/add", h.AddParticipant)
		participant.GET("/all", h.GetAllParticipants)
		participant.GET("/:id", h.GetParticipantById)
		participant.GET("/team/:id", h.GetParticipantsByTeamId)
		participant.PUT("/:id", h.UpdateParticipant)
		participant.DELETE("/:id", h.DeleteParticipantById)
		participant.GET("/user/:id", h.GetParticipantByCreatorId)
	}
	compare := v1.Group("/compare")
	{
		compare.POST("/two", h.compareTwoWithDescription)
		compare.POST("/lot", h.compareLot)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
