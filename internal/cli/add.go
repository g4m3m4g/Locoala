package cli

import (
	"fmt"
	"os"

	"github.com/g4m3m4g/locoala/internal/config"
	"github.com/g4m3m4g/locoala/internal/hosts"
)

func add() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: lc add <name> <port>")
		return
	}

	name := os.Args[2]
	port := os.Args[3]

	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	domain := name + ".local"
	target := "http://localhost:" + port

	cfg.Domains[domain] = target

	err = config.Save(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = hosts.Add(domain)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	
	fmt.Printf("Added %s → %s\n", domain, target)
}