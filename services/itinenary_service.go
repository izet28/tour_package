package services

import (
	"github.com/izet28/tour_package/models"
	"github.com/izet28/tour_package/repositories"
	"gorm.io/gorm"
)

// TourPackageService handles the business logic for tour packages
type ItineraryService struct {
	Repo repositories.ItineraryRepository
}

// NewItineraryService creates a new ItineraryService
func NewItineraryService(db *gorm.DB) ItineraryService {
	return ItineraryService{
		Repo: repositories.NewItineraryRepository(db),
	}
}

func (s *ItineraryService) GetAll() ([]models.Itinerary, error) {
	return s.Repo.GetAll()
}

func (s *ItineraryService) GetByID(id int) (*models.Itinerary, error) {
	return s.Repo.GetByID(id)
}

func (s *ItineraryService) Create(paket models.Itinerary) (*models.Itinerary, error) {
	return s.Repo.Create(paket)
}

func (s *ItineraryService) Update(id int, paket models.Itinerary) (*models.Itinerary, error) {
	return s.Repo.Update(id, paket)
}

func (s *ItineraryService) Delete(id int) error {
	return s.Repo.Delete(id)
}
