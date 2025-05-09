package main

import (
	"log/slog"
	"os"
	"time"
)

func UTCLogger(source bool) *slog.Logger {
	l := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: source,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{
					Key:   slog.TimeKey,
					Value: slog.TimeValue(time.Now().UTC()),
				}
			}
			return a
		},
	})

	return slog.New(l)
}