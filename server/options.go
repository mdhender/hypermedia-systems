// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package server

import (
	"log"
	"net/http"
)

type Option func(*Server) error

func WithApplication(h http.Handler) Option {
	return func(s *Server) error {
		s.Handler = h
		return nil
	}
}

func WithHost(host string) Option {
	return func(s *Server) error {
		s.host = host
		log.Printf("[server] set host to %q\n", s.host)
		return nil
	}
}

func WithPort(port string) Option {
	return func(s *Server) error {
		s.port = port
		log.Printf("[server] set port to %q\n", s.port)
		return nil
	}
}
