package server

import (
	"github.com/bicycledays/bridge/src/service"
	"log"
	"net/http"
)

type Server struct {
	Port        string
	Comparators map[string]service.Comparator
}

func (s *Server) Run() error {
	http.HandleFunc(RouteSockets, wsEndpoint)
	http.HandleFunc(RouteScanCom, checkComPorts)

	log.Println("start server")
	return http.ListenAndServe(":"+s.Port, nil)
}
