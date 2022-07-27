package log

import (
	"fmt"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func Info(format string, a ...any) {
	logger.Info(fmt.Sprintf(format, a))
}
