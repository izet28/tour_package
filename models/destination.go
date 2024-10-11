package models

// Destination defines the structure for a destination
type Destination struct {
	ID           uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string        `gorm:"size:100;not null" json:"name"   `              // Nama destinasi (misalnya, Jogja, Bali)
	TourPackages []TourPackage `gorm:"foreignKey:DestinationID" json:"tour_packages"` // Relasi one-to-many ke TourPackage
}
