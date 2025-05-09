package main

import (
	
)

var logger = UTCLogger(false)

func main() {
	logger.Info("This is an info message", "rotation", 34, "altitude", 500, "temperature", 64)
}