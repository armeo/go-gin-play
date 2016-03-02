package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestV1(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	expected := "pong"
	req, _ := http.NewRequest("GET", "/v1/ping", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/v1/ping", pongV1)
	r.ServeHTTP(w, req)

	if w.Body.String() != expected {
		t.Errorf("expect %v but %s", expected, w.Body.String())
	}
}

func TestV2(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	expected := `{"message":"pongpong"}`
	req, _ := http.NewRequest("GET", "/v2/ping", nil)
	w := httptest.NewRecorder()

	r := gin.Default()
	r.GET("/v2/ping", pongV2)
	r.ServeHTTP(w, req)

	if w.Body.String() != expected {
		t.Errorf("expect %v but %s", expected, w.Body.String())
	}
}
