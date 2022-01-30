package handler

import (
	"encoding/json"
	"github.com/bicycledays/bridge/src/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Sockets struct {
	Route   string
	Clients []websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (h *Handler) measure(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Ошибка при запуске сокетов", err.Error())
		return
	}

	log.Println("Client Connected", c.ClientIP())
	err = ws.WriteMessage(1, []byte("OK"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Ошибка при ответе клиенту", err.Error())
		return
	}

	_, message, _ := ws.ReadMessage()
	var comparator service.Comparator
	err = json.Unmarshal(message, &comparator)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Ошибка парсинга параметров компаратора", err.Error())
		return
	}
	_, err = h.service.CheckComparator(&comparator)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Параметры компаратора не валидны", err.Error())
		return
	}

	var measure []byte
	port := comparator.OpenPort()
	ch := make(chan []byte)
	go comparator.Listen(ch, port)
	go func() {
		tick := time.NewTicker(time.Millisecond * 500)
		defer tick.Stop()
		for {

		}
	}()

	for {
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
		measure = <-ch
		if err := ws.WriteMessage(websocket.TextMessage, measure); err != nil {
			newErrorResponse(
				c,
				http.StatusInternalServerError,
				"Ошибка при передаче сообщения по сокетам",
				err.Error(),
			)
			return
		}
	}
}
