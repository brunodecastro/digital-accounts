package server

import (
	"fmt"
	"github.com/brunodecastro/digital-accounts/app/config"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Server struct {
	Server http.Server
	Router *httprouter.Router
}

func NewServer() *Server {
	server := &Server{
		Server: http.Server{},
		Router: httprouter.New(),
	}

	// Set the api routes
	server.setRoutes()

	return server
}

func indexPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func healthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func (server *Server) setRoutes() {
	router := server.Router
	router.GET("/", indexPage)
	router.GET("/health-check", healthCheck)
	router.GET("/account", healthCheck)
}

func (server *Server) ListenAndServe(webServerConfig *config.WebServerConfig) error {
	return http.ListenAndServe(webServerConfig.GetWebServerAddress(), server.Router)
}