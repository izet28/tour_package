package models

import (
	"time"
)

// Itinerary defines the structure for itinerary items
type Itinerary struct {
	ID            uint ` gorm:"primaryKey;autoIncrement" json:"id"`
	TourPackageID uint ` gorm:"not null" json:"tour_package_id"`
	// TourPackage   TourPackage `gorm:"foreignKey:TourPackageID" json:"tourpackage" `             // Foreign key untuk TourPackage
	Day         int       ` gorm:"not null" json:"day" validate:"required,gt=0"`            // Hari ke berapa dalam paket wisata
	Activity    string    ` gorm:"size:255;not null" json:"activity" validate:"required" `  // Aktivitas yang dilakukan
	Description string    ` gorm:"type:text" json:"description" validate:"required,min=10"` // Deskripsi aktivitas
	StartTime   time.Time `json:"start_time" validate:"required" `                          // Waktu mulai aktivitas
	EndTime     time.Time `json:"end_time" validate:"required" `                            // Waktu akhir aktivitas

}
