package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/g4m3m4g/locoala/internal/config"
)

type Proxy struct {
	Config *config.Config
}

func (p *Proxy) Handler(w http.ResponseWriter, r *http.Request) {
	host := r.Host

	targetStr, ok := p.Config.Domains[host]
	if !ok {
		http.Error(w, "Domain not found", http.StatusNotFound)
		return
	}

	targetURL, err := url.Parse(targetStr)
	if err != nil {
		http.Error(w, "Invalid target", http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	log.Printf("%s → %s\n", host, targetStr)

	proxy.ServeHTTP(w, r)
}