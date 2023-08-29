// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"github.com/mdhender/hypermedia-systems/server"
	"github.com/spf13/cobra"
	"log"
)

// cmdServe is the root module for all server commands
var cmdServe = &cobra.Command{
	Use:   "serve",
	Short: "serve an application",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Printf("[serve] running persistentPreRun\n")
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

		s, err := server.New(options...)
		if err != nil {
			log.Fatal(err)
		}
		s.Serve()
	},
}

var argsServe = struct {
	host                       string
	port                       string
	doNotUseBadRunesMiddleware bool
	doNotUseCorsMiddleware     bool
}{
	port: "8080",
}
