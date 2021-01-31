package main

import (
	"fmt"
	"github.com/brunodecastro/digital-accounts/app/api/controller"
	"github.com/brunodecastro/digital-accounts/app/api/server"
	"github.com/brunodecastro/digital-accounts/app/common/logger"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/brunodecastro/digital-accounts/app/persistence/database/postgres"
	"github.com/brunodecastro/digital-accounts/app/service"
	"github.com/brunodecastro/digital-accounts/app/util"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	// Initialize app configs
	apiConfig := config.LoadConfigs()

	// Initialize app log implementation
	logger.InitLog(apiConfig)
	logger.LogApp.Info("Starting Digital Accounts API...")

	// Configure database pool connection
	databaseConnection := postgres.ConnectPoolConfig(&apiConfig.DatabaseConfig)
	defer databaseConnection.Close()

	err := postgres.UpMigrations(&apiConfig.DatabaseConfig)
	util.MaybeFatal(err, "Unable to execute postgres migrations.")

	server := createServer(databaseConnection)
	logger.LogApp.Info(fmt.Sprintf("Server running on %s ...", apiConfig.WebServerConfig.GetWebServerAddress()))
	err = server.ListenAndServe(&apiConfig.WebServerConfig)
	util.MaybeFatal(err, "Unable to start the web server.")
}

func createServer(databaseConnection *pgxpool.Pool) *server.Server {
	// Account services
	accountRepository := postgres.NewAccountRepository(databaseConnection)
	accountService := service.NewAccountService(accountRepository)
	accountController := controller.NewAccountController(accountService)

	// Authenticate services
	authenticateService := service.NewAuthenticateService(accountRepository)
	authenticateController := controller.NewAuthenticateController(authenticateService)

	// Transfer services
	transferRepository := postgres.NewTransferRepository(databaseConnection)
	transferService := service.NewTransferService(transferRepository, accountRepository)
	transferController := controller.NewTransferController(transferService)

	// Configure the webserver and serve
	return server.NewServer(
		authenticateController,
		accountController,
		transferController,
	)
}
