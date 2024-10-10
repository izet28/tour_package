package repositories

import (
	"github.com/izet28/tour_package/models"

	"gorm.io/gorm"
)

type DestinationRepository struct {
	DB *gorm.DB
}

func NewDestinationRepository(db *gorm.DB) DestinationRepository {
	return DestinationRepository{DB: db}
}

func (r *DestinationRepository) GetAll() ([]models.Destination, error) {
	var destinations []models.Destination
	err := r.DB.Find(&destinations).Error
	return destinations, err
}

func (r *DestinationRepository) GetByID(id int) (*models.Destination, error) {
	var destination models.Destination
	err := r.DB.First(&destination, id).Error
	return &destination, err
}

func (r *DestinationRepository) Create(destination models.Destination) (*models.Destination, error) {
	err := r.DB.Create(&destination).Error
	return &destination, err
}

func (r *DestinationRepository) Update(id int, updatedPackage models.Destination) (*models.Destination, error) {
	var destination models.Destination
	if err := r.DB.First(&destination, id).Error; err != nil {
		return nil, err
	}
	r.DB.Model(&destination).Updates(updatedPackage)
	return &destination, nil
}

func (r *DestinationRepository) Delete(id int) error {
	return r.DB.Delete(&models.Destination{}, id).Error
}
