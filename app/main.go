package app

import (
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/logger"
	"github.com/brunodecastro/digital-accounts/app/persistence"
)

func main() {
	// Initialize app configs
	var apiConfig = config.LoadConfigs()

	// Initialize app log implementation
	logger.InitLogFacade(apiConfig)

	logger.LogApp.Info("Starting Digital Accounts API...")

	// Configure database pool connection
	poolConfig := persistence.PoolConfig(&apiConfig.DatabaseConfig)
	defer poolConfig.Close()


}
