package server

import (
	"log"
	"net/http"

	"github.com/g4m3m4g/locoala/internal/proxy"
)

func Start(port int, p *proxy.Proxy) {
	http.HandleFunc("/", p.Handler)

	addr := ":80"

	log.Println("Proxy running on", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}