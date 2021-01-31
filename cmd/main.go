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
	logger.GetLogger().Info("Starting Digital Accounts API...")

	// Configure database pool connection
	databaseConnection := postgres.ConnectPoolConfig(&apiConfig.DatabaseConfig)
	defer databaseConnection.Close()

	err := postgres.UpMigrations(&apiConfig.DatabaseConfig)
	util.MaybeFatal(err, "Unable to execute postgres migrations.")

	server := createServer(databaseConnection, apiConfig)
	logger.GetLogger().Info(fmt.Sprintf("Server running on %s ...", apiConfig.WebServerConfig.GetWebServerAddress()))
	err = server.ListenAndServe(&apiConfig.WebServerConfig)
	util.MaybeFatal(err, "Unable to start the web server.")
}

func createServer(databaseConnection *pgxpool.Pool, apiConfig *config.Config) *server.Server {
	transactionHelper := postgres.NewTransactionHelper(databaseConnection)

	// Account services
	accountRepository := postgres.NewAccountRepository(databaseConnection, transactionHelper)
	accountService := service.NewAccountService(accountRepository, transactionHelper)
	accountController := controller.NewAccountController(accountService)

	// Authentication services
	authenticationService := service.NewAuthenticationService(accountRepository)
	authenticationController := controller.NewAuthenticationController(authenticationService)

	// Transfer services
	transferRepository := postgres.NewTransferRepository(databaseConnection, transactionHelper)
	transferService := service.NewTransferService(transferRepository, accountRepository, transactionHelper)
	transferController := controller.NewTransferController(transferService)

	// Configure the webserver and serve
	return server.NewServer(
		authenticationController,
		accountController,
		transferController,
	)
}
