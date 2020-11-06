package model

// Role 角色模型
type Role struct {
	BaseModel
	Alias string `gorm:"not null comment:'名称'"`
}

// RoleRouter 角色路由关联
type RoleRouter struct {
	BaseModel
	RoleID   int `gorm:"not null comment:'角色ID'"`
	RouterID int `gorm:"not null comment:'路由ID'"`
}
