// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"github.com/mdhender/hypermedia-systems/app/flask"
	"github.com/mdhender/hypermedia-systems/server"
	"github.com/spf13/cobra"
	"log"
)

// cmdServeFlask serves the sample flask
var cmdServeFlask = &cobra.Command{
	Use:   "flask",
	Short: "serve the sample flask",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Printf("[flask] running preRun\n")
	},
	Run: func(cmd *cobra.Command, args []string) {
		var options []server.Option
		if argsServe.host != "" {
			options = append(options, server.WithHost(argsServe.host))
		}
		if argsServe.port == "" {
			options = append(options, server.WithPort(argsServe.host))
		} else {
			options = append(options, server.WithPort("8080"))
		}
		if !argsServe.doNotUseBadRunesMiddleware {
			options = append(options, server.WithBadRunesMiddleware())
		}
		if !argsServe.doNotUseCorsMiddleware {
			options = append(options, server.WithCorsMiddleware())
		}

		a, err := flask.New()
		if err != nil {
			log.Fatal(err)
		}
		options = append(options, server.WithApplication(a))

		s, err := server.New(options...)
		if err != nil {
			log.Fatal(err)
		}
		s.Serve()
	},
}

var argsServeFlask struct{}
