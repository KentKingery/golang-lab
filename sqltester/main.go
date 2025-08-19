package main

func main() {
	logger.Debug("Startup")
	GetPartitions()
	logger.Debug("Shutdown")
}
