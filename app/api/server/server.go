package server

import (
	"fmt"
	"github.com/brunodecastro/digital-accounts/app/api/controller"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Server struct {
	Server            http.Server
	Router            *httprouter.Router
	accountController controller.AccountController
}

func NewServer(accountController controller.AccountController) *Server {
	server := &Server{
		Server:            http.Server{},
		Router:            httprouter.New(),
		accountController: accountController,
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
	router.GET("/accounts", server.accountController.GetAll)
	router.GET("/account/:account_id/balance", server.accountController.GetBalance)
}

func (server *Server) ListenAndServe(webServerConfig *config.WebServerConfig) error {
	return http.ListenAndServe(webServerConfig.GetWebServerAddress(), server.Router)
}
