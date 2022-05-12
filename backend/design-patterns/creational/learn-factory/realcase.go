package main

import (
	"net/http"
	"net/http/httptest"
)

// factory mock

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// type HttpRequest struct {
// 	httpReq *http.Client
// }

func NewHttpRequest() Doer {
	// httpReq := &HttpRequest{}
	return &http.Client{}
}

type MockHttpRequest struct{}

func (m *MockHttpRequest) Do(req *http.Request) (*http.Response, error) {
	res := httptest.NewRecorder()
	return res.Result(), nil
}

func NewMockHttpRequest() Doer {
	return &MockHttpRequest{}
}

func init() {

	httpReq := NewHttpRequest()
	httpReq.Do(nil)

	mockHttpReq := NewMockHttpRequest()
	mockHttpReq.Do(nil)
}
