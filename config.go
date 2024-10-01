package core

import (
	"github.com/Cellularhacker/logger-go"
	"github.com/joho/godotenv"
	"os"
	"time"
)

// MARK: For Program's running mode settings
const (
	ValueEnvModeProduction  = "production"
	ValueEnvModeDevelopment = "development"

	KeyEnvMode = "MODE"
)

var envMode = ValueEnvModeDevelopment

// MARK: For Program's localization settings
const (
	ValueEnvAsiaSeoul = "Asia/Seoul"
	ValueEnvAsiaTokyo = "Asia/Tokyo"

	KeyEnvLocation = "LOCATION"
)

var loc *time.Location

// MARK: init()
func init() {
	// MARK: <START>Initializing Variables
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	} else {
		logger.Info("Loaded environments from env file")
	}
	// MARK: <EMD>Initializing Variables

	// MARK: Set Program's running mode
	envMode = os.Getenv(KeyEnvMode)
	// MARK: Set Program's localization
	envLocation := os.Getenv(KeyEnvLocation)
	if len(envLocation) <= 0 {
		loc = time.Local
	} else {
		loc, err = time.LoadLocation(envLocation)
		if err != nil {
			logger.Fatal("github.com/Cellularhacker/core-go", "config.go", "init()", "time.LoadLocation(envLocation)", KeyEnvLocation, envLocation, err)
		}
	}
}

func SetLoc(tz *time.Location) {
	loc = tz
}
