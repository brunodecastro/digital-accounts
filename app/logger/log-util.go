package logger

import (
	"github.com/sirupsen/logrus"
)

// LogUtil enforces specific log message formats
type LogUtil struct {
	*logrus.Logger
}

func NewLogUtil() *LogUtil {
	var baseLogger = logrus.New()
	var logUtil = &LogUtil{baseLogger}
	return logUtil
}

// FatalError generates a standard fatal error
func (logUtil *LogUtil) FatalError(errorMessage string, err error) {
	logUtil.WithError(err).Fatal(errorMessage)
}
