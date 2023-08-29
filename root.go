// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"github.com/spf13/cobra"
	"log"
	"time"
)

// executeCobra initializes the Cobra tree and executes the root module.
// it returns some errors (mostly with command line arguments),
// but panics more often than not.
func executeCobra() error {
	defer func(started time.Time) {
		elapsed := time.Now().Sub(started)
		if argsRoot.timeSelf {
			log.Printf("elapsed time: %v\n", elapsed)
		}
	}(time.Now())

	cmdRoot.PersistentFlags().BoolVar(&argsRoot.timeSelf, "time", false, "time commands")

	cmdRoot.AddCommand(cmdVersion)

	cmdRoot.AddCommand(cmdServe)
	cmdServe.PersistentFlags().StringVar(&argsServe.host, "host", argsServe.host, "host to bind to")
	cmdServe.PersistentFlags().StringVar(&argsServe.port, "port", argsServe.port, "port to listen on")

	cmdServe.AddCommand(cmdServeFlask)

	return cmdRoot.Execute()
}

// cmdRoot represents the base command when called without any subcommands
var cmdRoot = &cobra.Command{
	Short:   "htmx: sample applications",
	Long:    `Sample applications from the Hypermedia Systems book.`,
	Version: "0.0.0",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		log.Printf("[root] running preRunE\n")
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if argsRoot.timeSelf {
			defer func(started time.Time) {
				elapsed := time.Now().Sub(started)
				log.Printf("elapsed time: %v\n", elapsed)
			}(time.Now())
		}

		return nil
	},
}

var argsRoot struct {
	timeSelf bool
}
