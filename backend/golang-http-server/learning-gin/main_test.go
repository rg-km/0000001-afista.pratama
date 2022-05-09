package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
)

func TestRouter(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/hello-name/afis", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	//assert.Equal(t, "pong", w.Body)
}
