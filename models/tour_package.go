package models

import (
	"time"
)

// TourPackage defines the structure for a tour package
type TourPackage struct {
	ID              uint      `    gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string    `    gorm:"size:100;not null" json:"name"`
	Description     string    `  gorm:"type:text" json:"description"`
	Price           float64   `      gorm:"not null" json:"price"`
	DurationDays    int       ` gorm:"not null" json:"duration_days"`
	StartDate       time.Time `  gorm:"not null" json:"start_date"`
	EndDate         time.Time `    gorm:"not null" json:"end_date"`
	MaxParticipants int       `gorm:"not null" json:"max_participants"`
	Availability    bool      ` gorm:"default:true" json:"availability"`

	DestinationID uint `gorm:"not null" json:"destination_id" ` // Foreign key ke Destination
	// Destination   Destination `gorm:"foreignKey:DestinationID" json:"destination" ` // Relasi ke Destination

	Itineraries []Itinerary `gorm:"foreignKey:TourPackageID" json:"itineraries"  ` // Relasi one-to-many ke Itinerary

}
