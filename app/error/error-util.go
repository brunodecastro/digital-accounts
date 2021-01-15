package error

import "github.com/brunodecastro/digital-accounts/app/logger"

func MaybeFatal(err error, errorMessage string) {
	if err != nil {
		logger.NewLogUtil().FatalError(errorMessage, err)
	}
}
