package repositories

import (
	"github.com/izet28/tour_package/models"

	"gorm.io/gorm"
)

type TncRepository struct {
	DB *gorm.DB
}

func NewTncRepository(db *gorm.DB) TncRepository {
	return TncRepository{DB: db}
}

func (r *TncRepository) GetAll() ([]models.Tnc, error) {
	var tncs []models.Tnc
	err := r.DB.Find(&tncs).Error
	return tncs, err
}

func (r *TncRepository) GetByID(id int) (*models.Tnc, error) {
	var tnc models.Tnc
	err := r.DB.First(&tnc, id).Error
	return &tnc, err
}

func (r *TncRepository) Create(tnc models.Tnc) (*models.Tnc, error) {
	err := r.DB.Create(&tnc).Error
	return &tnc, err
}

func (r *TncRepository) Update(id int, updatedPackage models.Tnc) (*models.Tnc, error) {
	var tnc models.Tnc
	if err := r.DB.First(&tnc, id).Error; err != nil {
		return nil, err
	}
	r.DB.Model(&tnc).Updates(updatedPackage)
	return &tnc, nil
}

func (r *TncRepository) Delete(id int) error {
	return r.DB.Delete(&models.Tnc{}, id).Error
}
