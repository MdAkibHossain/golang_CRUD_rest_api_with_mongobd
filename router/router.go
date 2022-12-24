package router

import (
	"golang_CRUD_rest_api_with_mongodb/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/get_all_movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/create_movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/mark_as_watched/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/delete_movie/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/delete_all_movies", controller.DeleteAllMovies).Methods("DELETE")
	return router
}
