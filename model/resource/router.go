package resource

import "goa/model"

// Router 路由模型
type Router struct {
	model.BaseModel
	Path     string `gorm:"not null"`
	Method   string `gorm:"not null"`
	ModuleID uint
}
