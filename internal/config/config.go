// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package config defines the configuration for the application.
package config

import (
	"flag"
	"fmt"
	"github.com/mdhender/hypermedia-systems/internal/homedir"
	"github.com/mdhender/hypermedia-systems/internal/semver"
	"path/filepath"
)

type Config struct {
	Action struct {
		Version bool
	}
	Home       string
	Middleware struct {
		BadRunes bool
		Cors     bool
		Logging  bool
	}
	Public string // path to public assets
	Semver semver.Version
	Server struct {
		Host string
		Port string
	}
	Templates string // path to template files
}

func New() (*Config, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, fmt.Errorf("home: %w", err)
	}

	c := &Config{
		Home:   home,
		Semver: semver.Version{Major: 0, Minor: 1},
	}
	c.Server.Port = "8080"

	flag.BoolVar(&c.Action.Version, "version", c.Action.Version, "show version")

	flag.StringVar(&c.Home, "home", c.Home, "path to home directory")
	flag.BoolVar(&c.Middleware.BadRunes, "bad-runes", c.Middleware.BadRunes, "enable bad-runes middleware")
	flag.BoolVar(&c.Middleware.Cors, "cors", c.Middleware.Cors, "enable cors middleware")
	flag.BoolVar(&c.Middleware.Logging, "logging", c.Middleware.Logging, "enable logging middleware")
	flag.StringVar(&c.Public, "public", c.Public, "path to public assets")
	flag.StringVar(&c.Server.Host, "host", c.Server.Host, "host to bind to")
	flag.StringVar(&c.Server.Port, "port", c.Server.Port, "port to listen on")
	flag.StringVar(&c.Templates, "templates", c.Templates, "path to template files")

	flag.Parse()

	if path, err := filepath.Abs(c.Public); err != nil {
		return nil, fmt.Errorf("public: %w", err)
	} else {
		c.Public = path
	}

	if path, err := filepath.Abs(c.Templates); err != nil {
		return nil, fmt.Errorf("templates: %w", err)
	} else {
		c.Templates = path
	}

	return c, nil
}
