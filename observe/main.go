package main

import (
	"log/slog"
)

func main() {
	slog.Debug("Debug")
	slog.Info("Info")
	slog.Error("Error")
	slog.Warn("Warn")
}
