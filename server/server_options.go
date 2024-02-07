package server

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// New will create an instance of new http server
func New(options ...func(*Server)) *Server {
	s := &Server{
		port:    ":8080",
		timeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	s.context = ctx
	s.cancelFunc = stop

	for _, o := range options {
		o(s)
	}

	return s
}

// WithRouter to assign router
func WithRouter(r http.Handler) func(*Server) {
	return func(s *Server) {
		s.router = r
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
