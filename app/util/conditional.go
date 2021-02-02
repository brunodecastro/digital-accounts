package util

import (
	"github.com/brunodecastro/digital-accounts/app/common/logger"
	"go.uber.org/zap"
)

// MaybeFatal - can give a fatal log depending on the condition
func MaybeFatal(err error, errorMessage string) {
	if err != nil {
		logger.GetLogger().Fatal(errorMessage, zap.Error(err))
	}
}

// MaybeError - can give a error log depending on the condition
func MaybeError(err error, errorMessage string) {
	if err != nil {
		logger.GetLogger().Fatal(errorMessage, zap.Error(err))
	}
}
