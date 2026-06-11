package server

import "time"

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

type Option func(*Server)

func WithHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func New(options ...Option) *Server {
	svr := &Server{}

	for _, o := range options {
		o(svr)
	}

	return svr
}

func (s *Server) Start() error {
	return nil
}
