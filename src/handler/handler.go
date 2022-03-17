package handler

import (
	"github.com/bicycledays/bridge/src/service"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"time"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.GET("/scan-com-ports", h.listPorts)
	router.GET("/ping", h.ping)
	router.GET("/measure", h.measure)

	api := router.Group("/api", h.comparatorIdentity)
	{
		api.POST("/print", h.print)
		api.POST("/tare", h.tare)
		api.POST("/f2", h.f2)
		api.POST("/f5", h.f5)
		api.POST("/f6", h.f6)
	}

	return router
}
