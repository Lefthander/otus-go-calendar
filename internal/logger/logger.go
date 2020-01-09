package logger

import (
	"go.uber.org/zap"
)

// InitLogger does initilization of sugared logger
func InitLogger() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	suggarLogger := logger.Sugar()
	return suggarLogger, nil
}
