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
	router.HandleFunc("/api/v2/auth/user/Register", handler.SignupV2()).Methods("POST")

	// Employee routes
	Employee := router.PathPrefix("/api/v1/employee").Subrouter()
	Employee.Use(auth.AuthMiddleware)
	Employee.Use(auth.RequireRole("Employee"))
	Employee.HandleFunc("/myDashBoard", handler.GetMyDashboard()).Methods("GET")

	admin := router.PathPrefix("/api/v1/admin").Subrouter()
	admin.Use(auth.AuthMiddleware)
	admin.Use(auth.RequireRole("admin"))
	admin.HandleFunc("/auth/registerByAdmin", handler.RegisterUserByAdmin()).Methods("POST")

	EmpManager := router.PathPrefix("/api/v1/EmpManager").Subrouter()
	EmpManager.Use(auth.AuthMiddleware)
	EmpManager.Use(auth.RequireRole("EmployeeManager"))

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
	AssetManager.HandleFunc("/Employee/RetriveAsset", handler.RetrieveAsset()).Methods("POST")
	AssetManager.HandleFunc("/DeleteAsset/{Asset_id}", handler.DeleteAsset()).Methods("DELETE") //make sure ki asset assign kisi ko na ho
	//// dynamic assing asset api
	AssetManager.HandleFunc("/Employee/AssignAsset", handler.DynamicAssignAssetHandler()).Methods("POST")
	// protected api not for Employee
	Protected := router.PathPrefix("/api/v1/Protected").Subrouter()
	Protected.Use(auth.AuthMiddleware)
	Protected.HandleFunc("/getAllAssets", handler.GetAllAssets()).Methods("GET")
	Protected.HandleFunc("/getAllEmployee", handler.GetAllEmployees()).Methods("GET")
	Protected.HandleFunc("/getAssetInfo/{Asset_id}", handler.GetAssetInfoHandler()).Methods("GET")
	Protected.HandleFunc("/getAssetDashBoard", handler.GetDashboard()).Methods("GET")
	Protected.HandleFunc("/SearchByName", handler.EmployeeSearchByName()).Methods("GET")
	Protected.HandleFunc("/SearchByEmail", handler.EmployeeSearchByEmail()).Methods("GET")
	Protected.HandleFunc("/SearchByPhoneNo", handler.EmployeeSearchByPhoneNo()).Methods("GET")
	Protected.HandleFunc("/AssetTimeLine/{Employee_Id}", handler.GETAssetTimeLine()).Methods("GET")
	Protected.HandleFunc("/GetAssignedList", handler.AssetAssignedStatus()).Methods("GET")
	Protected.HandleFunc("/GetUnAssignedList", handler.AssetUnAssignedStatus()).Methods("GET")
	Protected.HandleFunc("/ChangeRole", handler.ChangeRole()).Methods("PATCH")
	Protected.HandleFunc("/DeleteEmployee", handler.DeleteEmployee()).Methods("DELETE")

	return router
}
