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

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/scan-com-ports", h.listPorts)
	router.GET("/measure", h.measure)

	api := router.Group("/api", h.comparatorIdentity)
	{
		api.GET("/print", h.print)
		api.GET("/tare", h.tare)
		api.GET("/f2", h.f2)
		api.GET("/platform", h.platform)
	}

	return router
}
