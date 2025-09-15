package main

import (
	"github.com/matheusjv11/go-banking/app"
	"github.com/matheusjv11/go-banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
