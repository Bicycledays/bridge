package handler

import (
	"encoding/json"
	"github.com/bicycledays/bridge/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
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

	api := router.Group("/api", h.comparatorIdentity)
	{
		api.GET("/print", h.print)
		api.GET("/tare", h.tare)
	}

	return router
}

func tare(w http.ResponseWriter, r *http.Request) {
	var c service.Comparator
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func checkComPorts(w http.ResponseWriter, r *http.Request) {

}
