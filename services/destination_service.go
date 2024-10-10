package services

import (
	"github.com/izet28/tour_package/models"
	"github.com/izet28/tour_package/repositories"
	"gorm.io/gorm"
)

// TourPackageService handles the business logic for tour packages
type DestinationService struct {
	Repo repositories.DestinationRepository
}

// NewDestinationService creates a new DestinationService
func NewDestinationService(db *gorm.DB) DestinationService {
	return DestinationService{
		Repo: repositories.NewDestinationRepository(db),
	}
}

func (s *DestinationService) GetAll() ([]models.Destination, error) {
	return s.Repo.GetAll()
}

func (s *DestinationService) GetByID(id int) (*models.Destination, error) {
	return s.Repo.GetByID(id)
}

func (s *DestinationService) Create(paket models.Destination) (*models.Destination, error) {
	return s.Repo.Create(paket)
}

func (s *DestinationService) Update(id int, paket models.Destination) (*models.Destination, error) {
	return s.Repo.Update(id, paket)
}

func (s *DestinationService) Delete(id int) error {
	return s.Repo.Delete(id)
}
