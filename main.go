// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package main implements the applications from Hypermedia Systems.
// (See https://hypermedia.systems/ for the original source.)
package main

import (
	"github.com/mdhender/hypermedia-systems/pkg/cli"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.LUTC)

	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
