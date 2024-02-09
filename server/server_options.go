package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// New will create an instance of new http server
func New(router http.Handler, options ...func(*Server)) *Server {
	s := &Server{
		context:           context.Background(),
		port:              ":8080",
		router:            router,
		readHeaderTimeout: 3 * time.Second,
		readTimeout:       5 * time.Second,
		writeTimeout:      8 * time.Second,
		timeout:           10 * time.Second,
	}

	for _, o := range options {
		o(s)
	}

	return s
}

// WithPort to assign port
func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.port = fmt.Sprintf(":%d", port)
	}
}

// WithContext to assign context
func WithContext(ctx context.Context) func(*Server) {
	return func(s *Server) {
		s.context = ctx
	}
}

// WithTimeout for timeout
func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}
