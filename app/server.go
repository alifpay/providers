package app

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/alifpay/providers/models"
)

//Server - http api server
type Server struct {
	srv *http.Server //http server
	cfg models.Config
}

//New - new http api server
func New(config models.Config) *Server {
	config.Version = strings.ReplaceAll(config.Version, "/", "")
	config.Version = "/" + config.Version
	return &Server{cfg: config}
}

//Run http api server
func (s *Server) Run() error {

	s.srv = &http.Server{
		Addr:              s.cfg.Host,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      40 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       2 * time.Minute,
		Handler:           s.routers(),
	}

	err := s.srv.ListenAndServe()
	return err
}

//Shutdown - graceful shutdown of http api server
func (s *Server) Shutdown(ctx context.Context) {
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
