package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jajang57/gsk_dev/app/controllers"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFilehandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public").Handler(staticFilehandler).Methods("GET")

	staticFileDirectoryview := http.Dir("./view/")
	staticFilehandlerview := http.StripPrefix("/page/", http.FileServer(staticFileDirectoryview))
	server.Router.PathPrefix("/page").Handler(staticFilehandlerview).Methods("GET")

}
