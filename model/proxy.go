package model

// 代理主机模型
type ProxyHost struct {
	BaseModel
	URLScheme string `gorm:"not null comment:'目标协议'"`
	URLHost   string `gorm:"not null comment:'目标地址'"`
	URLPath   string `gorm:"not null comment:'目标路径'"`
	Host      string `gorm:"not null comment:'头地址'"`
}

// Proxy 代理资源模型
type Proxy struct {
	BaseModel
	Path     string `gorm:"not null comment:'路由路径'"`
	Method   string `gorm:"not null comment:'路由方法'"`
	ModuleID int    `gorm:"not null comment:'模块ID'"`
	HostID   int    `gorm:"not null comment:'主机ID'"`
}
