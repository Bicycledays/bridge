package handler

import (
	"github.com/bicycledays/bridge/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) print(c *gin.Context) {
	portName, _ := c.Get(comparatorCtx)
	comparator := h.service.Comparators[portName.(string)]
	port := comparator.OpenPort()

	ch := make(chan string)
	go comparator.Listen(ch, port)
	comparator.Send(port, service.Print)
	measure := <-ch
	c.JSON(http.StatusOK, map[string]string{
		"measure": measure,
	})
}

func (h *Handler) tare(c *gin.Context) {
	portName, _ := c.Get(comparatorCtx)
	comparator := h.service.Comparators[portName.(string)]
	port := comparator.OpenPort()

	ch := make(chan string)
	go comparator.Listen(ch, port)
	comparator.Send(port, service.Tare)
	measure := <-ch
	c.JSON(http.StatusOK, map[string]string{
		"measure": measure,
	})
}
