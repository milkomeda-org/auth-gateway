package model

// Proxy 代理模型
type Proxy struct {
	BaseModel
	Path     string `gorm:"not null comment:'路由路径'"`
	Method   string `gorm:"not null comment:'路由方法'"`
	ModuleID int    `gorm:"not null comment:'模块ID'"`
}
