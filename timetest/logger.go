package main

import (
	"log/slog"
	"os"
)

var opts = &slog.HandlerOptions{
    // Use the ReplaceAttr function on the handler options
    // to be able to replace any single attribute in the log output
    ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
         // check that we are handling the time key
         if a.Key != slog.TimeKey {
         	return a
         }
         
         t := a.Value.Time().UTC()
         
         // change the value from a time.Time to a String
         // where the string has the correct time format.
         a.Value = slog.StringValue(t.Format("2006-01-02T15:04:05.000Z"))
         
         return a
    },
}

var logger = slog.New(slog.NewJSONHandler(os.Stdout, opts))