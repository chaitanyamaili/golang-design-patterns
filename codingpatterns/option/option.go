package option

import (
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
}

type ServerOption func(*Server)

func NewServer(opts ...ServerOption) *Server {
	s := &Server{
		host:    "localhost",
		port:    8080,
		timeout: time.Second * 5,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func WitHost(host string) ServerOption {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}
