// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package cli

import (
	"github.com/mdhender/hypermedia-systems/app/flask"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

var (
	argsServeFlask = struct {
		Templates string // path to templates
	}{
		Templates: ".",
	}

	cmdServeFlask = &cobra.Command{
		Use:   "flask",
		Short: "run flask server",
		Long:  `Run an flask server.`,
		Run: func(cmd *cobra.Command, args []string) {
			path, err := filepath.Abs(argsServeFlask.Templates)
			if err != nil {
				log.Fatalf("[flask] templates: %v\n", err)
			}
			argsServeFlask.Templates = path

			log.Printf("[flask] serve flask %+v\n", argsServeFlask)

			var flaskOptions []flask.Option
			flaskOptions = append(flaskOptions, flask.WithTemplates(argsServeFlask.Templates))
			flaskOptions = append(flaskOptions, flask.WithContacts(flask.NewContacts(
				flask.NewContact("jim", "Jim"),
				flask.NewContact("joe", "Joe"),
			)))

			app, err := flask.New(flaskOptions...)
			if err != nil {
				log.Fatalf("[flask] app: %v\n", err)
			}
			argsServe.Handler = app

			cmdServe.Run(cmd, args)
		},
	}
)
