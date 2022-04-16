package main

import (
	"context"
	"flag"
	"log"

	"github.com/HotPotatoC/twitter-clone/user/clients"
	"github.com/HotPotatoC/twitter-clone/user/config"
)

var (
	configPath *string
)

func init() {
	// Config file path flag
	configPath = flag.String("config", "", "Path to config file (.yaml)")

	flag.Parse()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf, err := config.ReadYAMLFile(*configPath)
	if err != nil {
		// TODO: better logging
		log.Fatal(err)
	}

	clients, err := clients.Init(ctx, conf.Clients)
	if err != nil {
		log.Fatal(err)
	}

	_ = clients
}
