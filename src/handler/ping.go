package handler

import "github.com/gin-gonic/gin"

func (h *Handler) ping(c *gin.Context) {
	newResultResponse(c, "pong")
}
