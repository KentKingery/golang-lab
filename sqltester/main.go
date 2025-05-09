package main

import (

)

func main() {
	logger.Debug("Startup")
	getSources()
	logger.Debug("Shutdown")
}