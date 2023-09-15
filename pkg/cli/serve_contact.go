// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package cli

import (
	"github.com/mdhender/hypermedia-systems/app/contact"
	"github.com/mdhender/hypermedia-systems/server"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

var (
	argsServeContact = struct {
		Templates string // path to templates
	}{
		Templates: ".",
	}

	cmdServeContact = &cobra.Command{
		Use:   "contact",
		Short: "run contact server",
		Long:  `Run the contact application.`,
		Run: func(cmd *cobra.Command, args []string) {
			path, err := filepath.Abs(argsServeContact.Templates)
			if err != nil {
				log.Fatalf("[contact] templates: %v\n", err)
			}
			argsServeContact.Templates = path

			log.Printf("[contact] serve contact %+v\n", argsServeContact)

			var contactOptions []contact.Option
			contactOptions = append(contactOptions, contact.WithTemplates(argsServeContact.Templates))
			contactOptions = append(contactOptions, contact.WithContacts(contact.NewContacts(
				contact.NewContact("jim", "Jim"),
				contact.NewContact("joe", "Joe"),
			)))

			app, err := contact.New(contactOptions...)
			if err != nil {
				log.Fatalf("[contact] app: %v\n", err)
			}

			log.Printf("[contact] serve %+v\n", argsServe)

			serverOptions := []server.Option{server.WithApplication(app)}
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
				log.Fatalf("[contact] %v", err)
			} else if s.Handler == nil {
				log.Fatalf("[contact] missing handler\n")
			}

			if err := s.Serve(); err != nil {
				log.Fatalf("[contact] %v", err)
			}
		},
	}
)
