package web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := NewHTTPServer()
	server.Get("/get", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("get hello world"))
	})
	server.Post("/post", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("post hello world"))
	})
	server.Start(":8080")
}
