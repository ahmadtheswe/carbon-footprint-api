package entity

import (
	"time"

	"github.com/google/uuid"
)

type EmissionFactor struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Unit        string    `gorm:"type:varchar(100);not null" json:"unit"`
	Category    string    `gorm:"type:varchar(100);not null" json:"category"`
	SubCategory string    `gorm:"type:varchar(100);not null" json:"sub_category"`
	Value       float64   `gorm:"type:float;not null" json:"value"`
	Source      string    `gorm:"type:varchar(255);not null" json:"source"`
	UpdatedAt   time.Time `gorm:"type:timestamp;not null" json:"updated_at"`
}
