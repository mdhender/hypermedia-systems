// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package server implements a web server.
package server

import (
	"context"
	"errors"
	"fmt"
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
	host, port            string
	useCorsMiddleware     bool
	useBadRunesMiddleware bool
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
	if s.port == ":" {
		return nil, fmt.Errorf("missing port")
	}
	s.Addr = net.JoinHostPort(s.host, s.port)

	return s, nil
}

func (s *Server) Set(option Option) error {
	return option(s)
}

func (s *Server) Serve() error {
	if s.Addr == ":" {
		return fmt.Errorf("missing port")
	} else if s.Handler == nil {
		panic("!")
		return fmt.Errorf("missing handler")
	}

	// inject middleware as requested
	if s.useBadRunesMiddleware {
		s.Handler = handleBadRunes(s.Handler)
	}
	if s.useCorsMiddleware {
		s.Handler = optionsCors(s.Handler)
	}

	// set up stuff so that we can gracefully shut down the server and application
	serverCh := make(chan struct{})
	go func() {
		log.Printf("[server] serving %q\n", s.Addr)
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("[server] exited with: %v", err)
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
	return fmt.Errorf("interrupted and stopped")
}
