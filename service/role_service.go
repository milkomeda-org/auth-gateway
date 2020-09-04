package service

import (
	"github.com/gin-gonic/gin"
	"singo/model"
)

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

type RoleDeleteService struct {
	Id uint `form:"id" json:"id" binding:"required"`
}

// DeleteRole 删除角色
func (r *RoleDeleteService) DeleteRole(c *gin.Context) bool {
	//TODO 未完成功能
	return nil == model.DB.Where("id = ?", r.Id).Delete(model.Role{}).Error
}
