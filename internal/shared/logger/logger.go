package logger

import (
	"go.uber.org/zap"
)

func Init() func() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("failed to create logger: " + err.Error())
	}
	zap.ReplaceGlobals(logger)
	return func() { logger.Sync() }
}
