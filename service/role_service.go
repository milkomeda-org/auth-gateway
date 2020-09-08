package service

import (
	"github.com/gin-gonic/gin"
	"goa/model"
)

// RoleCreateService 角色创建服务
type RoleCreateService struct {
	Alias string `form:"alias" json:"alias" binding:"required"`
}

// CreateRole 创建角色
func (r *RoleCreateService) CreateRole(c *gin.Context) bool {
	role := model.Role{
		Alias: r.Alias,
	}
	return nil == model.DB.Create(&role).Error
}

// RoleDeleteService 角色删除服务
type RoleDeleteService struct {
	ID uint `form:"id" json:"id" binding:"required"`
}

// DeleteRole 删除角色
func (r *RoleDeleteService) DeleteRole(c *gin.Context) bool {
	return nil == model.DB.Where("id = ?", r.ID).Delete(model.Role{}).Error
}
