package handler

import (
	"github.com/bicycledays/bridge/src/service"
	"github.com/gin-gonic/gin"
)

const (
	comparatorCtx = "Name"
)

func (h *Handler) comparatorIdentity(c *gin.Context) {
	var comparator service.Comparator
	if err := c.BindJSON(&comparator); err != nil {
		panic("Ошибка парсинга параметров компаратора")
	}
	portName := h.service.CheckComparator(&comparator)
	c.Set(comparatorCtx, portName)
}
