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
	}

	team := v1.Group("/team")
	{
		team.POST("/add", h.AddTeam)
		team.GET("/all", h.GetAllTeams)
		team.GET("/:id", h.GetTeamById)
		team.PUT("/:id", h.UpdateTeam)
	}

	participant := v1.Group("/participant")
	{
		participant.POST("/add", h.AddParticipant)
		participant.GET("/all", h.GetAllParticipants)
		participant.GET("/:id", h.GetParticipantById)
		participant.GET("/team/:team_id", h.GetParticipantsByTeamId)
		participant.PUT("/:id", h.UpdateParticipant)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
