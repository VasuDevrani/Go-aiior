package router

import (
	"mongo/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/home", controllers.HomePage).Methods("GET")
	router.HandleFunc("/api/movie", controllers.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteall", controllers.DeleteAllMovie).Methods("DELETE")

	return router
}
