package logger

import (
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/util/constants"
	"go.uber.org/zap"
	"log"
)

var LogApp *LogFacade

type LogFacade struct {
	logImpl *zap.Logger
}

func InitLog(config *config.Config) {
	var err error
	var logZap *zap.Logger

	if config.Profile == constants.ProfileProd {
		logZap, err = zap.NewProduction()
	} else {
		logZap, err = zap.NewDevelopment()
	}

	if err != nil {
		log.Fatalf("Unable to initialize logger implementation")
	}
	LogApp = &LogFacade{logImpl: logZap}
}

// GetLogImplementation returns the zap log implementation
func (logFacade *LogFacade) GetZapLogImplementation() *zap.Logger {
	return logFacade.logImpl
}

// Info generates a standard info log
func (logFacade *LogFacade) Info(infoMessage string) {
	logFacade.logImpl.Info(infoMessage)
}

// Error generates a standard info log
func (logFacade *LogFacade) Error(errorMessage string, err error) {
	logFacade.logImpl.Error(errorMessage, zap.Error(err))
}

// FatalError generates a standard fatal error log
func (logFacade *LogFacade) FatalError(errorMessage string, err error) {
	logFacade.logImpl.Fatal(errorMessage, zap.Error(err))
}
