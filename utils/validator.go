package utils

import (
	"github.com/izet28/tour_package/models"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateTourPackage(paket models.TourPackage) error {
	return validate.Struct(paket)
}
