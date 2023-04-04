package transport

import (
	"awesomeProject/config"
	"awesomeProject/transport/handlers"
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

type Server struct {
	HTTP   *echo.Echo
	h      *handlers.Handlers
	config *config.Config
}

func NewServer(h *handlers.Handlers, config *config.Config) *Server {
	return &Server{h: h, config: config}
}

func (s Server) Run(ctx context.Context) error {
	s.HTTP = echo.New()
	s.routes()
	go func() {
		err := s.HTTP.Start(s.config.Port)
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Shutdown failed")
		}
	}()
	<-ctx.Done()
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()
	err := s.HTTP.Shutdown(ctxShutdown)
	if err != nil {
		log.Fatalf("Shutdown failed")
	}
	return nil
}
