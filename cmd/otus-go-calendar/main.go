package main

import (
	"github.com/Lefthander/otus-go-calendar/internal/version"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.Infof("The application Calendar ver. %v.%v.%v.%v is starting...", version.Version, version.BuildNumber, version.CommitHash, version.BuildTime)
}
