package util

import (
	"github.com/brunodecastro/digital-accounts/app/common/logger"
	"go.uber.org/zap"
)

func MaybeFatal(err error, errorMessage string) {
	if err != nil {
		logger.GetLogger().Fatal(errorMessage, zap.Error(err))
	}
}

func MaybeError(err error, errorMessage string) {
	if err != nil {
		logger.GetLogger().Fatal(errorMessage, zap.Error(err))
	}
}
