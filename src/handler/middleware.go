package handler

import (
	"github.com/bicycledays/bridge/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	comparatorCtx = "Name"
)

func (h *Handler) comparatorIdentity(c *gin.Context) {
	var comparator service.Comparator
	if err := c.BindJSON(&comparator); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Ошибка парсинга параметров компаратора", err.Error())
		return
	}
	portName, err := h.service.CheckComparator(&comparator)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Ошибка проверки компаратора", err.Error())
		return
	}
	c.Set(comparatorCtx, portName)
}
