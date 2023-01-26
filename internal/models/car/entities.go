package car

import (
	"time"

	"gorm.io/gorm"
)

type Car struct {
	ID        string `json:"id,omitempty" gorm:"primaryKey"`
	Model     string `json:"model"`
	Brand     string `json:"brand"`
	Year      int32  `json:"year"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
