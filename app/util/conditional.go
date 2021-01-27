package util

import "github.com/brunodecastro/digital-accounts/app/logger"

func MaybeFatal(err error, errorMessage string) {
	if err != nil {
		logger.LogApp.FatalError(errorMessage, err)
	}
}
