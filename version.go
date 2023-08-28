// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package main

import (
	"github.com/spf13/cobra"
	"log"
)

// cmdVersion runs the version command
var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "display application version",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("htmx version %s\n", cmdRoot.Version)
	},
}
