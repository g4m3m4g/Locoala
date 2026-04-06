package cli

import (
	"fmt"
	"os"
)

func Run() {
	if len(os.Args) < 2 {
		help()
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "start":
		start()
	case "add":
		add()
	case "remove":
		remove()
	case "list":
		list()
	default:
		help()
	}
}

func help() {
	fmt.Println(`
Locoala commands:
  lc start             Start proxy server
  lc add <name> <port> Add domain
  lc remove <name>     Remove domain
  lc list              Show domains
`)
}