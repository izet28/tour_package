package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/izet28/tour_package/config"
	"github.com/izet28/tour_package/models"
	routes "github.com/izet28/tour_package/routers"
)

func main() {
	// Initialize database
	db, err := config.InitializeDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.TourPackage{}, &models.Itinerary{}, &models.Tnc{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi database: %v", err)
	}
	// Initialize router
	router := mux.NewRouter()

	// Register routes
	routes.RegisterRoutes(router, db)

	// Start server
	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
