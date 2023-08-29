// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package main implements the applications from Hypermedia Systems.
// (See https://hypermedia.systems/ for the original source.)
package main

import (
	"fmt"
	"github.com/mdhender/hypermedia-systems/app/flask"
	"github.com/mdhender/hypermedia-systems/server"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.LUTC)
	log.Printf("[main] starting...\n")

	if err := run(); err != nil {
		log.Fatal(err)
	}

	log.Printf("[main] completed\n")
}

func run() error {
	root := &RootCommand{}
	return root.Execute(os.Args[1:])
}

type RootCommand struct {
	timeSelf bool
}

func (cmd *RootCommand) Execute(args []string) error {
	for len(args) != 0 {
		var arg string
		arg, args = args[0], args[1:]
		log.Printf("[rootCommand] %s %v\n", arg, args)
		if arg == "-h" || arg == "--help" {
			// do something
		} else if arg == "--time" {
			cmd.timeSelf = true
		} else if arg == "serve" {
			sub := &ServeCommand{}
			return sub.Execute(args)
		} else {
			return fmt.Errorf("unknown option %q", arg)
		}
	}
	if len(args) != 0 {
		return fmt.Errorf("unknown option %q", args[0])
	}
	return nil
}

type ServeCommand struct{}

func (cmd *ServeCommand) Execute(args []string) error {
	var options []server.Option
	options = append(options, server.WithPort("8080"))
	options = append(options, server.WithBadRunesMiddleware())
	options = append(options, server.WithCorsMiddleware())

	for len(args) != 0 {
		var arg string
		arg, args = args[0], args[1:]
		log.Printf("[serveCommand] %s %v\n", arg, args)
		if arg == "-h" || arg == "--help" {
			// do something
		} else if arg == "bad-runes-middleware" {
			options = append(options, server.WithBadRunesMiddleware())
		} else if arg == "--cors-middleware" {
			options = append(options, server.WithCorsMiddleware())
		} else if arg == "--host" {
			if len(args) == 0 {
				return fmt.Errorf("serve: --host requires hostname")
			}
			options = append(options, server.WithHost(args[0]))
			args = args[1:]
		} else if arg == "--no-bad-runes-middleware" {
			options = append(options, server.WithNoBadRunesMiddleware())
		} else if arg == "--no-cors-middleware" {
			options = append(options, server.WithNoCorsMiddleware())
		} else if arg == "--port" {
			if len(args) == 0 {
				return fmt.Errorf("serve: --port requires port value")
			}
			options = append(options, server.WithPort(args[0]))
			args = args[1:]
		} else if arg == "flask" {
			sub := ServeFlaskCommand{}
			app, err := sub.Execute(args)
			if err != nil {
				return err
			}
			options = append(options, server.WithApplication(app))
		} else {
			return fmt.Errorf("unknown option %q", arg)
		}
	}
	if len(args) != 0 {
		return fmt.Errorf("unknown option %q", args[0])
	}

	s, err := server.New(options...)
	if err != nil {
		log.Fatal(err)
	}
	return s.Serve()
}

type ServeFlaskCommand struct{}

func (cmd *ServeFlaskCommand) Execute(args []string) (http.Handler, error) {
	a, err := flask.New()
	if err != nil {
		log.Fatal(err)
	}

	for len(args) != 0 {
		var arg string
		arg, args = args[0], args[1:]
		log.Printf("[serveCommand] %s %v\n", arg, args)
		if arg == "-h" || arg == "--help" {
			// do something
		} else {
			return nil, fmt.Errorf("unknown option %q", arg)
		}
	}

	if len(args) != 0 {
		return nil, fmt.Errorf("unknown option %q", args[0])
	}

	return a, nil
}
