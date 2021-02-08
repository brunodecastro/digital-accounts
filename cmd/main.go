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
	"github.com/brunodecastro/digital-accounts/docs"
	"github.com/jackc/pgx/v4/pgxpool"

	_ "github.com/brunodecastro/digital-accounts/docs"
)

// @title Digital Accounts API
// @version 1.0
// @contact.name Bruno de Castro Oliveira
// @contact.email brunnodecastro@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	// Initialize app log implementation
	logger.GetLogger().Info("Starting Digital Accounts API...")

	// Swagger configs.
	docs.SwaggerInfo.Host = config.GetAPIConfigs().WebServerConfig.GetWebServerAddress()
	docs.SwaggerInfo.BasePath = "/"

	// Configure database pool connection
	databaseConnection := postgres.ConnectPoolConfig()
	defer databaseConnection.Close()

	if config.GetAPIConfigs().ExecuteMigration {
		err := postgres.UpMigrations(
			config.GetAPIConfigs().DatabaseConfig.GetDatabaseURI(),
			config.GetAPIConfigs().MigrationPath,
		)
		util.MaybeFatal(err, "Unable to execute postgres migrations.")
	}

	server := createServer(databaseConnection)
	logger.GetLogger().Info(fmt.Sprintf("Server running on %s ...", config.GetAPIConfigs().WebServerConfig.GetWebServerAddress()))

	// Start server
	err := server.ListenAndServe()
	util.MaybeFatal(err, "Unable to start the web server.")
}

func createServer(databaseConnection *pgxpool.Pool) *server.Server {
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
