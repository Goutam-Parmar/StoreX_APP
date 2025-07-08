package routes

import (
	"net/http"

	"StoreXApp/auth"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running"))
	}).Methods("GET")

	admin := router.PathPrefix("/api/v1/admin").Subrouter()

	admin.Use(auth.AuthMiddleware)
	admin.Use(auth.RequireRole("admin"))

	return router
}
