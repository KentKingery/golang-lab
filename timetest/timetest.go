package main

import (
	"os"
	"syscall"
	"os/signal"
	"time"
)

func main() {

	c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        logger.Error("Keyboard interrupt")
        os.Exit(1)
    }()
	
	for i := 1; i <= 100; i++ {
		logger.Info("System init",
		"iteration", i)
		time.Sleep(1 * time.Second)
	}

}