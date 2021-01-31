package util

import "github.com/brunodecastro/digital-accounts/app/common/logger"

func MaybeFatal(err error, errorMessage string) {
	if err != nil {
		logger.LogApp.FatalError(errorMessage, err)
	}
}

func MaybeError(err error, errorMessage string) {
	if err != nil {
		logger.LogApp.Error(errorMessage, err)
	}
}
