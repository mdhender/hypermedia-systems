// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package web2

import (
	"fmt"
	"github.com/mdhender/hypermedia-systems/internal/store"
	"net"
	"net/http"
	"os"
	"time"
)

type Option func(*App) error

func WithBadRunesMiddleware(handler func(next http.Handler) http.Handler) Option {
	return func(a *App) error {
		a.server.middleware.badRunes = handler
		return nil
	}
}

func WithCorsMiddleware(handler func(next http.Handler) http.Handler) Option {
	return func(a *App) error {
		a.server.middleware.cors = handler
		return nil
	}
}

func WithHost(host string) Option {
	return func(a *App) (err error) {
		a.server.host = host
		a.server.Addr = net.JoinHostPort(a.server.host, a.server.port)
		return nil
	}
}

func WithLoggingMiddleware(handler func(next http.Handler) http.Handler) Option {
	return func(a *App) error {
		a.server.middleware.logging = handler
		return nil
	}
}

func WithMaxBodyLength(l int) Option {
	return func(a *App) (err error) {
		a.server.MaxHeaderBytes = l
		return nil
	}
}
func WithPort(port string) Option {
	return func(a *App) (err error) {
		a.server.port = port
		a.server.Addr = net.JoinHostPort(a.server.host, a.server.port)
		return nil
	}
}

func WithReadTimeout(d time.Duration) Option {
	return func(a *App) error {
		a.server.ReadTimeout = d
		return nil
	}
}

func WithStore(s *store.Store) Option {
	return func(a *App) error {
		a.store = s
		return nil
	}
}

func WithTemplates(path string) Option {
	return func(a *App) error {
		if sb, err := os.Stat(path); err != nil {
			return err
		} else if !sb.IsDir() {
			return fmt.Errorf("invalid path %q", path)
		}
		a.templates = path
		return nil
	}
}

func WithWriteTimeout(d time.Duration) Option {
	return func(a *App) error {
		a.server.WriteTimeout = d
		return nil
	}
}
