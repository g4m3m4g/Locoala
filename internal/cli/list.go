package cli

import (
	"fmt"

	"github.com/g4m3m4g/locoala/internal/config"
)

func list() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("🌐 Domains:")

	for d, t := range cfg.Domains {
		fmt.Printf("  %s → %s\n", d, t)
	}
}