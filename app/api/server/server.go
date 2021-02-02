package server

import (
	"github.com/brunodecastro/digital-accounts/app/api/auth"
	"github.com/brunodecastro/digital-accounts/app/api/controller"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// Server - struct that represents the server api
type Server struct {
	Server                   http.Server
	Router                   *httprouter.Router
	authenticationController controller.AuthenticationController
	accountController        controller.AccountController
	transferController       controller.TransferController
}

// NewServer - server api instance
func NewServer(
	authenticationController controller.AuthenticationController,
	accountController controller.AccountController,
	transferController controller.TransferController) *Server {
	server := &Server{
		Server:                   http.Server{},
		Router:                   httprouter.New(),
		authenticationController: authenticationController,
		accountController:        accountController,
		transferController:       transferController,
	}

	// Config the api routes
	server.configRoutes()

	return server
}

// indexPage - index page off api
func indexPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Redirect to swagger ui page
	http.Redirect(w, r, "/swagger", http.StatusSeeOther)
}

// healthCheck - used to check if the app is alive
func healthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

// configRoutes - config the api routes
func (server *Server) configRoutes() {
	router := server.Router
	router.GET("/", indexPage)
	router.GET("/health-check", healthCheck)
	router.POST("/accounts", server.accountController.Create)
	router.GET("/accounts", server.accountController.FindAll)
	router.GET("/account/:account_id/balance", server.accountController.GetBalance)
	router.POST("/transfers", auth.AuthorizeMiddleware(server.transferController.Create))
	router.GET("/transfers", auth.AuthorizeMiddleware(server.transferController.FindAll))
	router.POST("/login", server.authenticationController.Authenticate)

	// Swagger UI
	router.HandlerFunc(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)
}

// ListenAndServe - listen and serve the api on host and port
func (server *Server) ListenAndServe(webServerConfig *config.WebServerConfig) error {
	return http.ListenAndServe(webServerConfig.GetWebServerAddress(), server.Router)
}
