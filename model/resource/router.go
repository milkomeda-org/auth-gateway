package resource

import "oa-auth/model"

// Router 路由模型
type Router struct {
	model.BaseModel
	Path     string `gorm:"not null comment:'路由路径'"`
	Method   string `gorm:"not null comment:'路由方法'"`
	ModuleID uint   `gorm:"not null comment:'模块ID'"`
}
