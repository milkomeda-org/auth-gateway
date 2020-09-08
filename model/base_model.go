package model

import "github.com/jinzhu/gorm"

// BaseModel model base
type BaseModel struct {
	gorm.Model
	Status int
}
