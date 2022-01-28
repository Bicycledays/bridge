package server

import (
	"encoding/json"
	"fmt"
	"github.com/bicycledays/bridge/src/service"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const RouteSockets = "/measure"

type Sockets struct {
	Route   string
	Clients []websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	for {
		log.Println("1")
		messageType, p, err := conn.ReadMessage()
		log.Print("messageType")
		log.Println(messageType)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		var c service.Comparator
		err = json.Unmarshal(p, &c)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println(*c.Config, *c.License)
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
		log.Println("4")
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println("end point")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}
	reader(ws)
}
