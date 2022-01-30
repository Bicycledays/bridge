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

	ch := make(chan []byte)
	go comparator.Listen(ch, port)
	err := comparator.Send(port, service.Print)
	if err != nil {
		newErrorResponse(
			c,
			http.StatusInternalServerError,
			"Ошибка при передаче команды на компаратор",
			err.Error(),
		)
		return
	}
	measure := <-ch
	newResultResponse(c, map[string]string{"measure": string(measure)})
}

func (h *Handler) tare(c *gin.Context) {
	portName, _ := c.Get(comparatorCtx)
	comparator := h.service.Comparators[portName.(string)]
	port := comparator.OpenPort()

	err := comparator.Send(port, service.Tare)
	if err != nil {
		newErrorResponse(
			c,
			http.StatusInternalServerError,
			"Ошибка при передаче команды на компаратор",
			err.Error(),
		)
		return
	}
	newResultResponse(c, nil)
}
