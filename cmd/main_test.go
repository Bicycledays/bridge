package main

import (
	"bytes"
	"encoding/json"
	"github.com/bicycledays/bridge/src/handler"
	"github.com/bicycledays/bridge/src/service"
	"github.com/gorilla/websocket"
	"github.com/tarm/serial"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()

	s := service.NewService()
	h := handler.NewHandler(s)
	return httptest.NewServer(h.InitRoutes())
}

func newTestWSServer(t *testing.T, route string) *websocket.Conn {
	t.Helper()

	ts := newTestServer(t)
	defer ts.Close()

	wsURL, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	wsURL.Scheme = "ws"
	u := wsURL.String() + route
	log.Println(u)

	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatal(err)
	}

	return ws
}

func newTestComparator(t *testing.T) *service.Comparator {
	t.Helper()

	return &service.Comparator{
		Config: &serial.Config{
			Name:     "/dev/ttyUSB0",
			Baud:     1200,
			Size:     7,
			Parity:   79,
			StopBits: 1,
		},
		Params: &service.Params{
			Id:     1,
			Number: "qwerty",
			Term:   "2022-02-05",
			Key:    "717765727479253125323032322d30322d3035",
		},
	}
}

func packComparatorToJson(t *testing.T) []byte {
	t.Helper()

	c := newTestComparator(t)
	js, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}

	return js
}

func checkResponse(t *testing.T, resp *http.Response) {
	t.Helper()

	// status
	if resp.StatusCode != 200 {
		t.Fatalf("status code: %v", resp.StatusCode)
	}

	// headers
	val, ok := resp.Header["Content-Type"]
	if !ok {
		t.Fatalf("expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("expected \"application/json; charset=utf-8\", got %s", val[0])
	}

	// body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read response body error: " + err.Error())
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		t.Fatalf("unmarshal body error: " + err.Error())
	}
	success, ok := result["success"]

	if !ok {
		t.Fatalf("body structure invalid: %s", string(bodyBytes))
	}
	switch v := success.(type) {
	case bool:
		if success != true {
			t.Fatalf("success false\nmessage: %s\nerror: %s", result["message"], result["error"])
		}
	default:
		t.Fatalf("success field invalid type: %T", v)
	}
}

func TestListPortsRoute(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/scan-com-ports")
	if err != nil {
		t.Fatalf("request error, got: %v", err)
	}
	checkResponse(t, resp)
}

func TestApiPrintRoute(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	js := packComparatorToJson(t)
	resp, err := http.Post(ts.URL+"/api/print", "application/json", bytes.NewBuffer(js))
	if err != nil {
		t.Fatalf("request error, got: %v", err)
	}
	checkResponse(t, resp)
}

func TestApiTareRoute(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	js := packComparatorToJson(t)
	resp, err := http.Post(ts.URL+"/api/tare", "application/json", bytes.NewBuffer(js))
	if err != nil {
		t.Fatalf("request error, got: %v", err)
	}
	checkResponse(t, resp)
}

func TestApiF2Route(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	js := packComparatorToJson(t)
	resp, err := http.Post(ts.URL+"/api/f2", "application/json", bytes.NewBuffer(js))
	if err != nil {
		t.Fatalf("request error, got: %v", err)
	}
	checkResponse(t, resp)
}

func TestApiPlatformRoute(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	js := packComparatorToJson(t)
	resp, err := http.Post(ts.URL+"/api/platform", "application/json", bytes.NewBuffer(js))
	if err != nil {
		t.Fatalf("request error, got: %v", err)
	}
	checkResponse(t, resp)
}

func TestMeasureRoute(t *testing.T) {
	ws := newTestWSServer(t, "/measure")
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			t.Fatal(err.Error())
		}
	}(ws)

	js := packComparatorToJson(t)
	if err := ws.WriteMessage(websocket.TextMessage, js); err != nil {
		t.Fatalf("%v", err)
	}
}
