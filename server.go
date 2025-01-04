package web

import (
	"fmt"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request)

type Server interface {
	http.Handler
	Start(addr string)
	AddRoute(method string, pattern string, handler Handler)
}

type HTTPServer struct {
	router *Router
}

// 确保 HTTPServer 肯定实现了 Server 接口
var _ Server = &HTTPServer{}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		router: NewRouter(),
	}
}

func (s *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler := s.router.Get(request.Method)
	if handler != nil {
		handler(writer, request)
	} else {
		str := fmt.Sprintf("未定义: method: %s, path %s", request.Method, request.URL.Path)
		writer.Write([]byte(str))
	}
}

func (s *HTTPServer) Start(addr string) {
	http.ListenAndServe(addr, s)
}

func (s *HTTPServer) AddRoute(method string, path string, handler Handler) {
	fmt.Printf("AddRoute --- method %s,  path %s \n", method, path)
	s.router.Add(method, handler)
}

func (s *HTTPServer) Get(path string, handler Handler) {
	s.AddRoute(http.MethodGet, path, handler)
}

func (s *HTTPServer) Post(path string, handler Handler) {
	s.AddRoute(http.MethodPost, path, handler)
}
