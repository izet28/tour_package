package models

// Itinerary defines the structure for itinerary items
type Itinerary struct {
	ID            uint ` gorm:"primaryKey;autoIncrement" json:"id"`
	TourPackageID uint ` gorm:"not null" json:"tour_package_id"` // Foreign key untuk TourPackage
	// Day            int            gorm:"not null" json:"day"                 // Hari ke berapa dalam paket wisata
	// Activity       string         gorm:"size:255;not null" json:"activity"   // Aktivitas yang dilakukan
	// Description    string         gorm:"type:text" json:"description"        // Deskripsi aktivitas
	// StartTime      time.Time      json:"start_time"                          // Waktu mulai aktivitas
	// EndTime        time.Time      json:"end_time"                            // Waktu akhir aktivitas
	// CreatedAt      time.Time      json:"created_at"
	// UpdatedAt      time.Time      json:"updated_at"
	// DeletedAt      gorm.DeletedAt gorm:"index" json:"-"
}
