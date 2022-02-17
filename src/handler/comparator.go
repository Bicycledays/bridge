package handler

import (
	"github.com/bicycledays/bridge/src/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) print(c *gin.Context) {
	portName, _ := c.Get(comparatorCtx)
	comparator := h.service.Comparators[portName.(string)]
	port, err := comparator.OpenPort()
	if err != nil {
		newErrorResponse(c, 400, "open serial port error", err.Error())
		return
	}
	comparator.Subscribers++
	log.Println("Subscribers++")
	log.Println(comparator.Subscribers)
	if comparator.Subscribers == 1 {
		go comparator.Listen(port)
		go comparator.SendWhileListing(port)
	}
	defer func() {
		comparator.Subscribers--
		log.Println("Subscribers--")
		log.Println(comparator.Subscribers)
	}()
	for comparator.Display == nil {
	}
	newResultResponse(c, map[string]string{"measure": string(comparator.Display)})
}

func (h *Handler) tare(c *gin.Context) {
	portName, _ := c.Get(comparatorCtx)
	comparator := h.service.Comparators[portName.(string)]
	port, err := comparator.OpenPort()
	if err != nil {
		newErrorResponse(c, 400, "open serial port error", err.Error())
		return
	}

	err = comparator.Send(port, service.Command{
		Format: 1,
		Symbol: service.Tare,
	})

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

func (h *Handler) f2(c *gin.Context) {
	portName, _ := c.Get(comparatorCtx)
	comparator := h.service.Comparators[portName.(string)]
	port, err := comparator.OpenPort()
	if err != nil {
		newErrorResponse(c, 400, "open serial port error", err.Error())
		return
	}

	err = comparator.Send(port, service.Command{
		Format: 2,
		Symbol: service.Func,
		Digit:  '2',
	})

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

func (h *Handler) f5(c *gin.Context) {
	portName, _ := c.Get(comparatorCtx)
	comparator := h.service.Comparators[portName.(string)]
	port, err := comparator.OpenPort()
	if err != nil {
		newErrorResponse(c, 400, "open serial port error", err.Error())
		return
	}

	err = comparator.Send(port, service.Command{
		Format: 2,
		Symbol: service.Func,
		Digit:  '5',
	})

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

func (h *Handler) f6(c *gin.Context) {
	portName, _ := c.Get(comparatorCtx)
	comparator := h.service.Comparators[portName.(string)]
	port, err := comparator.OpenPort()
	if err != nil {
		newErrorResponse(c, 400, "open serial port error", err.Error())
		return
	}

	err = comparator.Send(port, service.Command{
		Format: 2,
		Symbol: service.Func,
		Digit:  '6',
	})

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
