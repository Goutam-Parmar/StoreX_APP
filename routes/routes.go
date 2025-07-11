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
	// routes/routes.go (or wherever you init routes)
	router.HandleFunc("/api/v2/auth/user/Register", handler.SignupV2()).Methods("POST")

	admin := router.PathPrefix("/api/v1/admin").Subrouter()
	admin.Use(auth.AuthMiddleware)
	admin.Use(auth.RequireRole("admin"))
	admin.HandleFunc("/auth/registerByAdmin", handler.RegisterUserByAdmin()).Methods("POST")

	EmpManager := router.PathPrefix("/api/v1/EmpManager").Subrouter()
	EmpManager.Use(auth.AuthMiddleware)
	EmpManager.Use(auth.RequireRole("employeemanager"))

	EmpManager.HandleFunc("/auth/registerByEmpManager", handler.RegisterUserByEmpManager()).Methods("POST")

	AssetManager := router.PathPrefix("/api/v1/AssetManager").Subrouter()
	AssetManager.Use(auth.AuthMiddleware)
	AssetManager.Use(auth.RequireRole("AssetManager"))
	AssetManager.HandleFunc("/createLaptopAsset", handler.CreateLaptopAssetHandler()).Methods("POST")
	AssetManager.HandleFunc("/createMobileAsset", handler.CreateMobileAssetHandler()).Methods("POST")
	AssetManager.HandleFunc("/createMouseAsset", handler.CreateMouseAssetHandler()).Methods("POST")
	AssetManager.HandleFunc("/createMonitorAsset", handler.CreateMonitorAssetHandler()).Methods("POST")
	AssetManager.HandleFunc("/createHarddiskAsset", handler.CreateHarddiskAssetHandler()).Methods("POST")
	AssetManager.HandleFunc("/createPendriveAsset", handler.CreatePendriveAssetHandler()).Methods("POST")
	AssetManager.HandleFunc("/createAccessoriesAsset", handler.CreateAccessoriesAssetHandler()).Methods("POST")
	// asset distribution
	AssetManager.HandleFunc("/Employee/AssignLaptop", handler.LaptopAssignHandler()).Methods("POST")
	AssetManager.HandleFunc("/Employee/AssignMobile", handler.MobileAssignHandler()).Methods("POST")
	AssetManager.HandleFunc("/Employee/AssignMonitor", handler.MonitorAssignHandler()).Methods("POST")
	AssetManager.HandleFunc("/Employee/AssignMouse", handler.MouseAssignHandler()).Methods("POST")
	AssetManager.HandleFunc("/Employee/AssignHardDisc", handler.HardDiskAssignHandler()).Methods("POST")
	AssetManager.HandleFunc("/Employee/AssignPendrive", handler.PendriveAssignHandler()).Methods("POST")
	AssetManager.HandleFunc("/Employee/AssignAccessories", handler.AcessoriesAssignHandler()).Methods("POST")

	return router
}
