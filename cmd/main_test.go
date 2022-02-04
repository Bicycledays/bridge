package main

import (
	"encoding/json"
	"github.com/bicycledays/bridge/src/handler"
	"github.com/bicycledays/bridge/src/service"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestServer() *httptest.Server {
	s := service.NewService()
	h := handler.NewHandler(s)
	return httptest.NewServer(h.InitRoutes())
}

func checkResponse(t *testing.T, resp *http.Response) {
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
	ts := getTestServer()
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/scan-com-ports")
	if err != nil {
		t.Fatalf("request error, got: %v", err)
	}
	checkResponse(t, resp)
}
