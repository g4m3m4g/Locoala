package cli

import (
	"fmt"

	"github.com/g4m3m4g/locoala/internal/config"
	"github.com/g4m3m4g/locoala/internal/proxy"
	"github.com/g4m3m4g/locoala/internal/server"
)

func start() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	p := &proxy.Proxy{
		Config: cfg,
	}

	server.Start(cfg.Port, p)
}