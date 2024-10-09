package routes

import (
	"github.com/gorilla/mux"
	"github.com/izet28/tour_package/controllers"
	"gorm.io/gorm"
)

func RegisterRoutes(router *mux.Router, db *gorm.DB) {
	tourPackageController := controllers.NewTourPackageController(db)

	router.HandleFunc("/packages", tourPackageController.GetAll).Methods("GET")
	router.HandleFunc("/packages/{id}", tourPackageController.GetByID).Methods("GET")
	router.HandleFunc("/packages", tourPackageController.Create).Methods("POST")
	router.HandleFunc("/packages/{id}", tourPackageController.Update).Methods("PUT")
	router.HandleFunc("/packages/{id}", tourPackageController.Delete).Methods("DELETE")
}
