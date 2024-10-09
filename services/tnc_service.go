package services

import (
	"github.com/izet28/tour_package/models"
	"github.com/izet28/tour_package/repositories"
	"gorm.io/gorm"
)

// TourPackageService handles the business logic for tour packages
type TncService struct {
	Repo repositories.TncRepository
}

// NewTncService creates a new TncService
func NewTncService(db *gorm.DB) TncService {
	return TncService{
		Repo: repositories.NewTncRepository(db),
	}
}

func (s *TncService) GetAll() ([]models.Tnc, error) {
	return s.Repo.GetAll()
}

func (s *TncService) GetByID(id int) (*models.Tnc, error) {
	return s.Repo.GetByID(id)
}

func (s *TncService) Create(paket models.Tnc) (*models.Tnc, error) {
	return s.Repo.Create(paket)
}

func (s *TncService) Update(id int, paket models.Tnc) (*models.Tnc, error) {
	return s.Repo.Update(id, paket)
}

func (s *TncService) Delete(id int) error {
	return s.Repo.Delete(id)
}
