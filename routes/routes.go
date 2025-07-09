package routes

import (
	"StoreXApp/auth"
	"StoreXApp/handler"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	// Public route
	router.HandleFunc("/api/v1/auth/user/login", handler.Login()).Methods("POST")
	router.HandleFunc("/api/v1/auth/user/SelfRegister", handler.RegisterSelf()).Methods("POST")

	admin := router.PathPrefix("/api/v1/admin").Subrouter()
	admin.Use(auth.AuthMiddleware)
	admin.Use(auth.RequireRole("admin"))
	admin.HandleFunc("/auth/registerByAdmin", handler.RegisterUserByAdmin()).Methods("POST")

	EmpManager := router.PathPrefix("/api/v1/EmpManager").Subrouter()
	EmpManager.Use(auth.AuthMiddleware)
	EmpManager.Use(auth.RequireRole("EmpManager"))
	EmpManager.HandleFunc("/auth/registerByEmpManager", handler.RegisterUserByEmpManager()).Methods("POST")

	AssetManager := router.PathPrefix("/api/v1/AssetManager").Subrouter()
	AssetManager.Use(auth.AuthMiddleware)
	AssetManager.Use(auth.RequireRole("AssetManager"))
	AssetManager.HandleFunc("/createAsset", handler.CreateLaptopAssetHandler()).Methods("POST")

	return router
}
