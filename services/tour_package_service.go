package services

import (
	"github.com/izet28/tour_package/models"
	"github.com/izet28/tour_package/repositories"
	"gorm.io/gorm"
)

// TourPackageService handles the business logic for tour packages
type TourPackageService struct {
	Repo repositories.TourPackageRepository
}

// NewTourPackageService creates a new TourPackageService
func NewTourPackageService(db *gorm.DB) TourPackageService {
	return TourPackageService{
		Repo: repositories.NewTourPackageRepository(db),
	}
}

func (s *TourPackageService) GetAll() ([]models.TourPackage, error) {
	return s.Repo.GetAll()
}

func (s *TourPackageService) GetByID(id int) (*models.TourPackage, error) {
	return s.Repo.GetByID(id)
}

func (s *TourPackageService) Create(paket models.TourPackage) (*models.TourPackage, error) {
	return s.Repo.Create(paket)
}

func (s *TourPackageService) Update(id int, paket models.TourPackage) (*models.TourPackage, error) {
	return s.Repo.Update(id, paket)
}

func (s *TourPackageService) Delete(id int) error {
	return s.Repo.Delete(id)
}
