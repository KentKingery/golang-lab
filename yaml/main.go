package main

import (
	"log/slog"
	"os"

	"github.com/goccy/go-yaml"
)

var app struct {
	Name    string
	Version string
	Routes  []struct {
		Path string
		Url  string
	}
}

func main() {
	slog.Info("YAML Parser")
	yamlFile, err := os.ReadFile("test.yaml")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	if err := yaml.Unmarshal([]byte(yamlFile), &app); err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info(app.Routes[0].Path)
	slog.Info(app.Routes[0].Url)
}
