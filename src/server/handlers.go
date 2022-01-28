package server

import (
	"encoding/json"
	"github.com/bicycledays/bridge/src/service"
	"net/http"
)

const (
	RouteScanCom = "/scan-com-ports"
	RouteTare    = "tare"
	RoutePrint   = "tare"
)

func tare(w http.ResponseWriter, r *http.Request) {
	var c service.Comparator
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func checkComPorts(w http.ResponseWriter, r *http.Request) {

}
