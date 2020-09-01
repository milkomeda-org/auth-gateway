package model

import "github.com/jinzhu/gorm"

//base model
type BaseModel struct {
	gorm.Model
	Status int
}
