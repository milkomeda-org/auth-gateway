package model

//UserRoleMapping User和Role关系映射
type UserRoleMapping struct {
	BaseModel
	UserID uint
	RoleID uint
}
