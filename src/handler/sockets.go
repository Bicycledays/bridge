package handler

import (
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

func newErrorSocketResponse(ws *websocket.Conn, message string) {
	closeMessage := websocket.FormatCloseMessage(websocket.CloseNormalClosure, message)
	_ = ws.WriteMessage(websocket.CloseMessage, closeMessage)
	_ = ws.Close()
}

func (h *Handler) measure(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Ошибка при запуске сокетов", err.Error())
		return
	}

	log.Println("Client Connected", c.ClientIP())
	err = ws.WriteMessage(websocket.TextMessage, []byte("OK"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Ошибка при ответе клиенту", err.Error())
		return
	}

	_, message, _ := ws.ReadMessage()
	var portName string
	portName, err = h.service.CheckComparator(message)
	if err != nil {
		newErrorSocketResponse(ws, "comparator parameters are not valid:"+err.Error())
		return
	}
	comparator := h.service.Comparators[portName]
	err = ws.WriteMessage(websocket.TextMessage, []byte("OK"))
	port, err := comparator.OpenPort()
	if err != nil {
		newErrorSocketResponse(ws, "open serial port error:"+err.Error())
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

	ticker := time.NewTicker(time.Millisecond * 500)
	for {
		<-ticker.C
		if comparator.Display != nil {
			message = comparator.Display
			err = ws.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				newErrorSocketResponse(ws, "Ошибка при передаче сообщения по сокетам")
				return
			}
		}
	}
}
