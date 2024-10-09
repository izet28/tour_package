package models

import (
	"time"

	"gorm.io/gorm"
)

type TourPackage struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string         `gorm:"size:100;not null" json:"name"`
	Description     string         `gorm:"type:text" json:"description"`
	Destination     string         `gorm:"size:100;not null" json:"destination"`
	Price           float64        `gorm:"not null" json:"price"`
	DurationDays    int            ` gorm:"not null" json:"duration_days"`
	StartDate       time.Time      ` gorm:"not null" json:"start_date"`
	EndDate         time.Time      `gorm:"not null" json:"end_date"`
	MaxParticipants int            `gorm:"not null" json:"max_participants"`
	Availability    bool           `gorm:"default:true" json:"availability"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
