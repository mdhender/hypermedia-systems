// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package server implements a web server.
package server

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	http.Server
	host, port string
}

func New(options ...Option) (*Server, error) {
	// create a new http server with good values for timeouts and transports
	s := &Server{}
	s.IdleTimeout = 10 * time.Second
	s.ReadTimeout = 2 * time.Second
	s.WriteTimeout = 2 * time.Second

	// apply the options
	for _, option := range options {
		if err := option(s); err != nil {
			return nil, err
		}
	}
	s.Addr = net.JoinHostPort(s.host, s.port)

	return s, nil
}

func (s *Server) Serve() {
	if s.Addr == ":" {
		log.Fatalf("[serve] missing host and/or port\n")
	}

	// set up stuff so that we can gracefully shut down the server and application
	serverCh := make(chan struct{})
	go func() {
		log.Printf("[serve] serving %q\n", s.Addr)
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("[serve] exited with: %v", err)
		}
		close(serverCh)
	}()

	// create a catch for signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	// Wait for interrupt
	<-signalCh

	// use the context to shut down the application
	log.Printf("[server] received interrupt, shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("[server] failed to shutdown server: %s", err)
	}

	// If we got this far, it was an interrupt, so don't exit cleanly
	log.Fatalf("[server] interrupted and stopped\n")
}
