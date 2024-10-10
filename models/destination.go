package models

type Destination struct {
	ID          uint   ` gorm:"primaryKey;autoIncrement" json:"id"`
	Description string `gorm:"type:text" json:"description" validate:"required,min=3"`
}
