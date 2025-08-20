package main

import (
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/goccy/go-yaml"
	//"net/http"
	//"net/http/httputil"
	//"net/url"
)

type GwConfig struct {
	Version string
	Routes  []struct {
		Path string
		Url  string
	}
}

func main() {
	slog.Info("API gateway started")

	gwconfig := parseConfig("gwconfig.yaml")

	slog.Info(gwconfig.Version)
	for _, route := range gwconfig.Routes {
		slog.Info("Adding handler for " + route.Path)
		serviceUrl, _ := url.Parse(route.Url)
		proxy := httputil.NewSingleHostReverseProxy(serviceUrl)

		http.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			r.URL.Path = r.URL.Path[len(route.Path):] // Trim prefix for backend
			proxy.ServeHTTP(w, r)
		})
	}

	slog.Info("API gateway stopped")

}

func parseConfig(filePath string) GwConfig {
	var config GwConfig

	content, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("Error reading file: %v", err)
		os.Exit(1)
	}

	// Convert the byte slice to a string
	yamlString := string(content)

	if err := yaml.Unmarshal([]byte(yamlString), &config); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return config
}
