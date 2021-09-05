package webserver

import (
	"net/http"
)

type Server struct {
	Address string
	Handler http.Handler
}

func NewServer(addr string, handler http.Handler) (server Server, err error) {
	return Server{Address: addr, Handler: handler}, nil
}

func (s *Server) Run() {
	_ = http.ListenAndServe(s.Address, s.Handler)
}
