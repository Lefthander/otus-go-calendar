package main

import (
	"log"

	"github.com/Lefthander/otus-go-calendar/internal/logger"
	"github.com/Lefthander/otus-go-calendar/internal/version"
)

func main() {
	//	logger := logrus.New()
	logger, err := logger.InitLogger()
	if err != nil {
		log.Fatal("Unable to initialize logger", err)
	}
	logger.Infof("The application Calendar ver. %v.%v.%v.%v is starting...", version.Version, version.BuildNumber, version.CommitHash, version.BuildTime)
	logger.Debug("Debug output example...")
	logger.Error("Error output example")
	logger.Infof("Ending...")
}
