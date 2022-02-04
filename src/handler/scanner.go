package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) listPorts(c *gin.Context) {
	err := h.service.Scanner.RefreshPorts()
	if err != nil {
		newErrorResponse(c, 400, "refresh ports list error", err.Error())
	}
	newResultResponse(c, h.service.Scanner.GetPorts())
}
