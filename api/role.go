package api

import (
	"github.com/gin-gonic/gin"
	"singo/service"
)

// CreateRole 创建角色
func CreateRole(c *gin.Context) {
	var rs service.RoleCreateService
	if err := c.ShouldBind(&rs); err == nil {
		c.JSON(200, rs.CreateRole(c))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// DeleteRole 删除角色
func DeleteRole(c *gin.Context) {
	var rs service.RoleDeleteService
	if err := c.ShouldBind(&rs); err == nil {
		c.JSON(200, rs.DeleteRole(c))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
