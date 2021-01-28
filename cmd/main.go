package main

import (
	"fmt"
	"github.com/brunodecastro/digital-accounts/app/api/controller"
	"github.com/brunodecastro/digital-accounts/app/api/server"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/logger"
	"github.com/brunodecastro/digital-accounts/app/persistence/database/postgres"
	"github.com/brunodecastro/digital-accounts/app/service"
	"github.com/brunodecastro/digital-accounts/app/util"
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

	accountRepository := postgres.NewAccountRepository(databaseConnection)
	accountService := service.NewAccountService(accountRepository)
	accountController := controller.NewAccountController(accountService)

	err := postgres.UpMigrations(&apiConfig.DatabaseConfig)
	util.MaybeFatal(err, "Unable to execute postgres migrations.")

	// Configure the webserver and serve
	server := server.NewServer(accountController)
	logger.LogApp.Info(fmt.Sprintf("Server running on %s ...", apiConfig.WebServerConfig.GetWebServerAddress()))
	err = server.ListenAndServe(&apiConfig.WebServerConfig)
	util.MaybeFatal(err, "Unable to start the web server.")
}
