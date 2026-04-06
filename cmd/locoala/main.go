package main

import (
	"log"

	"github.com/g4m3m4g/locoala/internal/config"
	"github.com/g4m3m4g/locoala/internal/proxy"
	"github.com/g4m3m4g/locoala/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	p := &proxy.Proxy{
		Config: cfg,
	}

	server.Start(cfg.Port, p)
}