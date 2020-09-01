package model

//UserRole关系映射
type UserRoleMapping struct {
	BaseModel
	UserId uint
	RoleId uint
}
