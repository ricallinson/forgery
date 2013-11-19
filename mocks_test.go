package f

import (
	"errors"
	"net/http"
)

/*
   Create a Mock http.ResponseWriter for testing.
*/

type MockResponseWriter struct {
	error   bool
	headers http.Header
	Written []byte
}

func (this *MockResponseWriter) Header() http.Header {
	return this.headers
}

func (this *MockResponseWriter) Write(data []byte) (int, error) {
	if this.error {
		return 0, errors.New("")
	}
	this.Written = data
	return len(data), nil
}

func (this *MockResponseWriter) WriteHeader(code int) {
	return
}

func NewMockResponseWriter(error bool) *MockResponseWriter {
	return &MockResponseWriter{error: error, headers: make(http.Header)}
}

/*
   Create a Mock Renderer.
*/

type MockRenderer struct {
}

func (this *MockRenderer) Render(v string, i ...interface{}) (string, error) {
	return v, nil
}
