package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) listPorts(c *gin.Context) {
	h.service.Scanner.RefreshPorts()
	c.JSON(http.StatusOK, h.service.Scanner.GetPorts())
}
