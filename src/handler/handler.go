package handler

import (
	"github.com/bicycledays/bridge/src/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/scan-com-ports", h.listPorts)
	router.GET("/measure", h.measure)

	api := router.Group("/api", h.comparatorIdentity)
	{
		api.POST("/print", h.print)
		api.POST("/tare", h.tare)
		api.POST("/f2", h.f2)
		api.POST("/platform", h.platform)
	}

	return router
}
