package models

import (
	"time"

	"gorm.io/gorm"
)

type TourPackage struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id" validate:"required,min=3"`
	Name            string         `gorm:"size:100;not null" json:"name" validate:"required"`
	Description     string         `gorm:"type:text" json:"description" validate:"required,min=3"`
	Destination     string         `gorm:"size:100;not null" json:"destination" validate:"required,min=15"`
	Price           float64        `gorm:"not null" json:"price" validate:"required,gt=0"`
	DurationDays    int            ` gorm:"not null" json:"duration_days" validate:"required,gt=0"`
	StartDate       time.Time      ` gorm:"not null" json:"start_date" validate:"required"`
	EndDate         time.Time      `gorm:"not null" json:"end_date" validate:"required,gtefield=StartDate"`
	MaxParticipants int            `gorm:"not null" json:"max_participants" validate:"required,gt=0"`
	Availability    bool           `gorm:"default:true" json:"availability"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
