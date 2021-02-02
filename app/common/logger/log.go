package logger

import (
	"github.com/brunodecastro/digital-accounts/app/common/constants"
	"github.com/brunodecastro/digital-accounts/app/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"sync"
)

var (
	logImpl *zap.Logger
	doOnce  sync.Once
)

// GetLogger returns the zap log implementation
func GetLogger() *zap.Logger {
	doOnce.Do(func() {
		var err error
		var zapConfig zap.Config

		if config.GetAPIConfigs().Profile == constants.ProfileProd {
			zapConfig = zap.NewProductionConfig()
		} else {
			zapConfig = zap.NewDevelopmentConfig()
		}

		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		logImpl, err = zapConfig.Build(
			zap.AddCaller(),
			zap.AddStacktrace(zapcore.ErrorLevel),
		)
		if err != nil {
			log.Fatalf("Unable to initialize logger implementation")
		}
	})
	return logImpl
}
