package core

import (
	"github.com/Cellularhacker/logger"
	"github.com/joho/godotenv"
	"os"
)

const (
	ValueEnvModeProduction  = "production"
	ValueEnvModeDevelopment = "development"

	KeyEnvMode = "MODE"
)

var envMode = ValueEnvModeDevelopment

func init() {
	// MARK: <START>Initializing Variables
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	} else {
		logger.Info("Loaded environments from env file")
	}
	// MARK: <EMD>Initializing Variables

	envMode = os.Getenv(KeyEnvMode)
}
