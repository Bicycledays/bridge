package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	comparatorCtx = "Name"
)

func (h *Handler) comparatorIdentity(c *gin.Context) {
	js, err := c.GetRawData()
	log.Println(string(js))
	portName, err := h.service.CheckComparator(js)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Ошибка проверки компаратора", err.Error())
		return
	}
	c.Set(comparatorCtx, portName)
}
