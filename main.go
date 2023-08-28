// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package main implements the applications from Hypermedia Systems.
// (See https://hypermedia.systems/ for the original source.)
package main

import (
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.LUTC)
	log.Printf("[main] starting...\n")

	if err := run(); err != nil {
		log.Fatal(err)
	}

	log.Printf("[main] completed\n")
}
