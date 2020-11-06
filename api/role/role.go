package role

import (
	"oa-auth/serializer"
	"oa-auth/service/role"

	"github.com/gin-gonic/gin"
)

// Create 创建角色
func Create(c *gin.Context) {
	var rs role.RoleCreateService
	if err := c.ShouldBind(&rs); err == nil {
		c.JSON(200, rs.CreateRole(c))
	} else {
		c.JSON(200, serializer.I18Error(err))
	}
}

// Delete 删除角色
func Delete(c *gin.Context) {
	var rs role.RoleDeleteService
	if err := c.ShouldBind(&rs); err == nil {
		c.JSON(200, rs.DeleteRole(c))
	} else {
		c.JSON(200, serializer.I18Error(err))
	}
}
