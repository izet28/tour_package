package repositories

import (
	"github.com/izet28/tour_package/models"

	"gorm.io/gorm"
)

type TourPackageRepository struct {
	DB *gorm.DB
}

func NewTourPackageRepository(db *gorm.DB) TourPackageRepository {
	return TourPackageRepository{DB: db}
}

func (r *TourPackageRepository) GetAll() ([]models.TourPackage, error) {
	var packages []models.TourPackage
	err := r.DB.Find(&packages).Error
	return packages, err
}

func (r *TourPackageRepository) GetByID(id int) (*models.TourPackage, error) {
	var paket models.TourPackage
	err := r.DB.First(&paket, id).Error
	return &paket, err
}

func (r *TourPackageRepository) Create(paket models.TourPackage) (*models.TourPackage, error) {
	err := r.DB.Create(&paket).Error
	return &paket, err
}

func (r *TourPackageRepository) Update(id int, updatedPackage models.TourPackage) (*models.TourPackage, error) {
	var paket models.TourPackage
	if err := r.DB.First(&paket, id).Error; err != nil {
		return nil, err
	}
	r.DB.Model(&paket).Updates(updatedPackage)
	return &paket, nil
}

func (r *TourPackageRepository) Delete(id int) error {
	return r.DB.Delete(&models.TourPackage{}, id).Error
}
