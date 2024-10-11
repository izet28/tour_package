package routes

import (
	"github.com/gorilla/mux"
	"github.com/izet28/tour_package/controllers"
	"gorm.io/gorm"
)

func RegisterRoutes(router *mux.Router, db *gorm.DB) {
	tourPackageController := controllers.NewTourPackageController(db)
	tncController := controllers.NewTncController(db)
	destinationController := controllers.NewDestinationController(db)
	itineraryController := controllers.NewItineraryController(db)

	router.HandleFunc("/packages", tourPackageController.GetAll).Methods("GET")
	router.HandleFunc("/packages/{id}", tourPackageController.GetByID).Methods("GET")
	router.HandleFunc("/packages", tourPackageController.Create).Methods("POST")
	router.HandleFunc("/packages/{id}", tourPackageController.Update).Methods("PUT")
	router.HandleFunc("/packages/{id}", tourPackageController.Delete).Methods("DELETE")

	router.HandleFunc("/tnc", tncController.GetAll).Methods("GET")
	router.HandleFunc("/tnc/{id}", tncController.GetByID).Methods("GET")
	router.HandleFunc("/tnc", tncController.Create).Methods("POST")
	router.HandleFunc("/tnc/{id}", tncController.Update).Methods("PUT")
	router.HandleFunc("/tnc/{id}", tncController.Delete).Methods("DELETE")

	router.HandleFunc("/destination", destinationController.GetAll).Methods("GET")
	router.HandleFunc("/destination/{id}", destinationController.GetByID).Methods("GET")
	router.HandleFunc("/destination", destinationController.Create).Methods("POST")
	router.HandleFunc("/destination/{id}", destinationController.Update).Methods("PUT")
	router.HandleFunc("/destination/{id}", destinationController.Delete).Methods("DELETE")

	router.HandleFunc("/itinerary", itineraryController.GetAll).Methods("GET")
	router.HandleFunc("/itinerary/{id}", itineraryController.GetByID).Methods("GET")
	router.HandleFunc("/itinerary", itineraryController.Create).Methods("POST")
	router.HandleFunc("/itinerary/{id}", itineraryController.Update).Methods("PUT")
	router.HandleFunc("/itinerary/{id}", itineraryController.Delete).Methods("DELETE")
}
