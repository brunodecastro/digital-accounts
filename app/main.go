package main

import (
	"fmt"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/logger"
	"github.com/brunodecastro/digital-accounts/app/persistence/database/postgres"
	"github.com/brunodecastro/digital-accounts/app/util"
	"github.com/brunodecastro/digital-accounts/server"
)

func main() {
	// Initialize app configs
	apiConfig := config.LoadConfigs()

	// Initialize app log implementation
	logger.InitLogFacade(apiConfig)
	logger.LogApp.Info("Starting Digital Accounts API...")

	// Configure database pool connection
	databaseConnection := postgres.ConnectPoolConfig(&apiConfig.DatabaseConfig)
	defer databaseConnection.Close()

	//accountRepository := postgres.NewAccountRepository(poolConfig)
	//accountService := service.NewAccountService(accountRepository)

	err := postgres.UpMigrations(&apiConfig.DatabaseConfig)
	util.MaybeFatal(err, "Unable to execute postgres migrations.")

	// Configure the webserver and serve
	server := server.NewServer()
	logger.LogApp.Info(fmt.Sprintf("Server running on %s ...", apiConfig.WebServerConfig.GetWebServerAddress()))
	err = server.ListenAndServe(&apiConfig.WebServerConfig)
	util.MaybeFatal(err, "Unable to start the web server.")
}
