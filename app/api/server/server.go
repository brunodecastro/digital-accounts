package server

import (
	"fmt"
	"github.com/brunodecastro/digital-accounts/app/api/auth"
	"github.com/brunodecastro/digital-accounts/app/api/controller"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Server struct {
	Server                   http.Server
	Router                   *httprouter.Router
	authenticationController controller.AuthenticationController
	accountController        controller.AccountController
	transferController       controller.TransferController
}

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

	// Set the api routes
	server.setRoutes()

	return server
}

func indexPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Digital Accounts Api!\n")
}

func healthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func (server *Server) setRoutes() {
	router := server.Router
	router.GET("/", indexPage)
	router.GET("/health-check", healthCheck)
	router.POST("/accounts", server.accountController.Create)
	router.GET("/accounts", server.accountController.FindAll)
	router.GET("/account/:account_id/balance", server.accountController.GetBalance)
	router.POST("/transfers", auth.AuthorizeMiddleware(server.transferController.Create))
	router.GET("/transfers", server.transferController.FindAll)
	router.POST("/login", server.authenticationController.Authenticate)
}

func (server *Server) ListenAndServe(webServerConfig *config.WebServerConfig) error {
	return http.ListenAndServe(webServerConfig.GetWebServerAddress(), server.Router)
}
