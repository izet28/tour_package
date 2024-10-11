package utils

import (
	"github.com/izet28/tour_package/models"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateTourPackage(paket models.TourPackage) error {
	return validate.Struct(paket)
}

func ValidateTnc(tnc models.Tnc) error {
	return validate.Struct(tnc)
}
func ValidateDestination(destination models.Destination) error {
	return validate.Struct(destination)
}
func ValidateItinerary(itinerary models.Itinerary) error {
	return validate.Struct(itinerary)
}
