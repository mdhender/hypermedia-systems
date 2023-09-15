// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package cli

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
	"time"
)

func Execute() error {
	defer func(started time.Time) {
		if argsRoot.TimeSelf {
			log.Printf("[main] elapsed time %v\n", time.Now().Sub(started))
		}
	}(time.Now())

	if home, err := homedir.Dir(); err != nil {
		return err
	} else if argsRoot.Home, err = filepath.Abs(home); err != nil {
		return err
	}

	cmdRoot.PersistentFlags().BoolVar(&argsRoot.TimeSelf, "time", argsRoot.TimeSelf, "display run time statistics on completion")

	cmdServe.PersistentFlags().StringVar(&argsServe.Host, "host", argsServe.Host, "host to bind to")
	cmdServe.PersistentFlags().StringVar(&argsServe.Port, "port", argsServe.Port, "port to listen on")
	cmdServe.PersistentFlags().BoolVar(&argsServe.BadRunesMiddleware, "bad-runes-middleware", argsServe.BadRunesMiddleware, "enable bad runes middleware")
	cmdServe.PersistentFlags().BoolVar(&argsServe.CORSMiddleware, "cors-middleware", argsServe.CORSMiddleware, "enable CORS options middleware")
	cmdRoot.AddCommand(cmdServe)

	cmdServeContact.Flags().StringVar(&argsServeContact.Templates, "templates", argsServeContact.Templates, "path to templates")
	cmdServe.AddCommand(cmdServeContact)

	cmdRoot.AddCommand(cmdVersion)

	if err := cmdRoot.Execute(); err != nil {
		return err
	}

	return nil
}

var (
	// argsRoot is the global arguments
	argsRoot = struct {
		Home     string
		TimeSelf bool
	}{}

	// cmdRoot represents the base command when called without any subcommands
	cmdRoot = &cobra.Command{
		Use:   "htmx",
		Short: "An implementation of Hypermedia Systems",
		Long: `This application implements the example code from
the Hypermedia Systems book.`,
	}
)
