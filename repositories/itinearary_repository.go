package repositories

import (
	"github.com/izet28/tour_package/models"

	"gorm.io/gorm"
)

type ItineraryRepository struct {
	DB *gorm.DB
}

func NewItineraryRepository(db *gorm.DB) ItineraryRepository {
	return ItineraryRepository{DB: db}
}

func (r *ItineraryRepository) GetAll() ([]models.Itinerary, error) {
	var itinerary []models.Itinerary
	err := r.DB.Find(&itinerary).Error
	return itinerary, err
}

func (r *ItineraryRepository) GetByID(id int) (*models.Itinerary, error) {
	var itinerary models.Itinerary

	err := r.DB.First(&itinerary, id).Error
	return &itinerary, err
}

func (r *ItineraryRepository) Create(itinerary models.Itinerary) (*models.Itinerary, error) {
	err := r.DB.Create(&itinerary).Error
	return &itinerary, err
}

func (r *ItineraryRepository) Update(id int, updatedItinerary models.Itinerary) (*models.Itinerary, error) {
	var itinerary models.Itinerary
	if err := r.DB.First(&itinerary, id).Error; err != nil {
		return nil, err
	}
	r.DB.Model(&itinerary).Updates(updatedItinerary)
	return &itinerary, nil
}

func (r *ItineraryRepository) Delete(id int) error {
	return r.DB.Delete(&models.Itinerary{}, id).Error
}
