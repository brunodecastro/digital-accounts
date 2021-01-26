package main

import (
	"fmt"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/logger"
	"github.com/brunodecastro/digital-accounts/app/persistence"
	"github.com/brunodecastro/digital-accounts/app/util/conditional"
	"github.com/brunodecastro/digital-accounts/server"
)

func main() {
	// Initialize app configs
	apiConfig := config.LoadConfigs()

	// Initialize app log implementation
	logger.InitLogFacade(apiConfig)
	logger.LogApp.Info("Starting Digital Accounts API...")

	// Configure database pool connection
	poolConfig := persistence.PoolConfig(&apiConfig.DatabaseConfig)
	defer poolConfig.Close()

	// Configure the webserver and serve
	server := server.NewServer()
	logger.LogApp.Info(fmt.Sprintf("Server running on %s ...", apiConfig.WebServerConfig.GetWebServerAddress()))
	err := server.ListenAndServe(&apiConfig.WebServerConfig)
	conditional.MaybeFatal(err, "Unable to start the web server.")
}
