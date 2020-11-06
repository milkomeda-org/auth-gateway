package role

import (
	"github.com/gin-gonic/gin"
	"oa-auth/initializer/db"
	role2 "oa-auth/model/role"
)

// RoleCreateService 角色创建服务
type RoleCreateService struct {
	Alias string `form:"alias" json:"alias" binding:"required"`
}

// CreateRole 创建角色
func (r *RoleCreateService) CreateRole(c *gin.Context) bool {
	role := role2.Role{
		Alias: r.Alias,
	}
	return nil == db.DB.Create(&role).Error
}

// RoleDeleteService 角色删除服务
type RoleDeleteService struct {
	ID int `form:"id" json:"id" binding:"required"`
}

// DeleteRole 删除角色
func (r *RoleDeleteService) DeleteRole(c *gin.Context) bool {
	return nil == db.DB.Where("id = ?", r.ID).Delete(role2.Role{}).Error
}
