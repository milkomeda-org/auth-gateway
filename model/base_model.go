package model

import (
	"time"
)

// BaseModel model base
type BaseModel struct {
	ID        uint       `gorm:"primary_key;unique"`
	CreatedAt time.Time  `gorm:"type:datetime;not null"`
	UpdatedAt time.Time  `gorm:"type:datetime;not null"`
	DeletedAt *time.Time `sql:"index"`
	Status    int        `gorm:"not null;default:0"`
}
