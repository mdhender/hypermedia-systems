// Copyright (c) 2023 Michael D Henderson. All rights reserved.

// Package main implements the web1 contacts app from Hypermedia Systems.
// (See https://hypermedia.systems/ for the original source.)
package main

import (
	"fmt"
	"github.com/mdhender/hypermedia-systems/internal/config"
	"github.com/mdhender/hypermedia-systems/internal/server/web1"
	"log"
	"os"
	"time"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	} else if cfg.Action.Version {
		fmt.Printf("%s\n", cfg.Semver.String())
		os.Exit(0)
	}

	log.SetFlags(log.LstdFlags | log.LUTC)

	defer func(started time.Time) {
		log.Printf("[main] elapsed time %v\n", time.Now().Sub(started))
	}(time.Now())

	app, err := web1.New(cfg)
	if err != nil {
		log.Fatalf("[web1] app: %v\n", err)
	} else if err = app.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
