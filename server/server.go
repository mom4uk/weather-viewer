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

func (s *Server) InitStatic() {
	fs := http.FileServer(http.Dir("./static"))
	s.Mux.Handle("/static/", http.StripPrefix("/static/", fs))
}

func (s *Server) Start() error {
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	return http.ListenAndServe(addr, s.Mux)
}
