package cli

import (
	"fmt"
	"os"

	"github.com/g4m3m4g/locoala/internal/config"
	"github.com/g4m3m4g/locoala/internal/hosts"
)

func remove() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: lc remove <name>")
		return
	}

	name := os.Args[2]
	domain := name + ".local"

	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	delete(cfg.Domains, domain)

	err = config.Save(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = hosts.Remove(domain)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Removed %s\n", domain)
}