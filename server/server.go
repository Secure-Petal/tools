package server

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

// Server to manage server
// taken from https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go
type Server struct {
	port              string
	router            http.Handler
	context           context.Context
	timeout           time.Duration
	readHeaderTimeout time.Duration
	readTimeout       time.Duration
	writeTimeout      time.Duration
}

// Run to start the server
func (s *Server) Run() {
	ctx, stop := signal.NotifyContext(s.context, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:              s.port,
		Handler:           s.router,
		ReadHeaderTimeout: s.readHeaderTimeout,
		ReadTimeout:       s.readTimeout,
		WriteTimeout:      s.writeTimeout,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has a few seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
