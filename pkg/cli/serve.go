// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package cli

import (
	"github.com/mdhender/hypermedia-systems/server"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

var (
	argsServe = struct {
		Host               string
		Port               string
		BadRunesMiddleware bool
		CORSMiddleware     bool
		Handler            http.Handler
	}{
		Port: "8080",
	}

	cmdServe = &cobra.Command{
		Use:   "serve",
		Short: "run application server",
		Long:  `Run an application server.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("[flask] serve %+v\n", argsServe)

			if argsServe.Handler == nil {
				log.Fatalf("[serve] missing handler\n")
			}

			serverOptions := []server.Option{server.WithApplication(argsServe.Handler)}
			if argsServe.Host != "" {
				serverOptions = append(serverOptions, server.WithHost(argsServe.Host))
			}
			serverOptions = append(serverOptions, server.WithPort(argsServe.Port))
			if argsServe.BadRunesMiddleware {
				serverOptions = append(serverOptions, server.WithBadRunesMiddleware())
			}
			if argsServe.CORSMiddleware {
				serverOptions = append(serverOptions, server.WithCorsMiddleware())
			}
			s, err := server.New(serverOptions...)
			if err != nil {
				log.Fatalf("[serve] %v", err)
			} else if s.Handler == nil {
				log.Fatalf("[serve] missing handler\n")
			}

			if err := s.Serve(); err != nil {
				log.Fatalf("[serve] %v", err)
			}
		},
	}
)
