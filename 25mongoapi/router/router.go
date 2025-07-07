package router

import (
	"github.com/ffaiyaz23/golangLco/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movie", controller.GetMyAllMovie).Methods("GET")

	return router
}
