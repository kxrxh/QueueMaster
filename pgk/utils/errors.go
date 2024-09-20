package utils

import (
	"go.uber.org/zap"
)

// FailOnError logs the given error with the given message and exits the process.
// This should be used to report fatal errors in the application.
func FailOnError(err error, msg string) {
	if err != nil {
		logger, _ := zap.NewProduction()
		logger.Sugar().Fatalf("%s: %s", msg, err)
		logger.Sync()
	}
}
