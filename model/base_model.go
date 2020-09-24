package model

import (
	"time"
)

// BaseModel model base
type BaseModel struct {
	ID        uint       `gorm:"primary_key;unique" json:"id"`
	CreatedAt time.Time  `gorm:"type:datetime;not null" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:datetime;not null" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Status    int        `gorm:"not null;default:0" json:"status"`
}
