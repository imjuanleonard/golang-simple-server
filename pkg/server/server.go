package server

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/imjuanleonard/palu-covid/config"
	"github.com/imjuanleonard/palu-covid/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	srv *http.Server
}

func New(router *mux.Router) (*Server, error) {
	server := &Server{
		srv: &http.Server{
			Handler:      router,
			Addr:         config.Server.Address,
			ReadTimeout:  config.Server.ReadTimeoutSecond,
			WriteTimeout: config.Server.WriteTimeoutSecond,
			IdleTimeout:  config.Server.IdleTimeoutSecond,
		},
	}
	return server, nil
}

func (s Server) StartHTTPServer() {
	if s.srv != nil {
		go func() {
			err := s.srv.ListenAndServe()
			if err != http.ErrServerClosed {
				logger.Errorf("Error when starting the server = %v", err)
			}
		}()
		logger.Infof("Starting http server at - %s", s.srv.Addr)
		waitForShutdown(s.srv)
	}
}

func waitForShutdown(server *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)
	_ = <-sig
	logger.Infof("API server shutting down")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
	logger.Infof("API server shutdown complete")
}
