package server

import (
	"net/http"
	"os"
)

type Server struct {
	Mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		Mux: http.NewServeMux(),
	}
}

func (s *Server) GetMux() *http.ServeMux {
	return s.Mux
}

func (s *Server) Start() error {
	addr := os.Getenv("SERVER_ADDR")
	return http.ListenAndServe(addr, s.Mux)
}
