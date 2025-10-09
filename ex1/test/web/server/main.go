package server

import (
	"net/http"
)

type Server struct {
}

func (s Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}
