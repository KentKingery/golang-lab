package main

import (
	"log/slog"
	//"os"
	"strconv"
	"time"
)

func main() {
	year, month, day := time.Now().Date()
	slog.Info(strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day))
}
