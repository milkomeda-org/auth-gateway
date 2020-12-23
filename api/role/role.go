package role

import (
	"auth-gateway/serializer"
	"auth-gateway/service/role"

	"github.com/gin-gonic/gin"
)

// Create 创建角色
func Create(c *gin.Context) *serializer.Response {
	var rs role.CreateService
	if err := c.ShouldBind(&rs); err != nil {
		return serializer.I18Error(err)
	}
	if e := rs.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// Delete 删除角色
func Delete(c *gin.Context) *serializer.Response {
	var rs role.DeleteService
	if err := c.ShouldBind(&rs); err != nil {
		return serializer.I18Error(err)
	}
	if e := rs.Execute(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// View 查看角色
func View(c *gin.Context) *serializer.Response {
	var rs role.ViewService
	if err := c.ShouldBind(&rs); err != nil {
		return serializer.I18Error(err)
	}
	v, e := rs.Execute()
	if nil != e {
		return serializer.Failed(e)
	}
	return serializer.Success(v)
}

// ModuleAdd 添加模块
func ModuleAdd(c *gin.Context) *serializer.Response {
	var os role.ModuleService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Add(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// ModuleRemove 移除模块
func ModuleRemove(c *gin.Context) *serializer.Response {
	var os role.ModuleService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Remove(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// ProxyAdd 添加代理
func ProxyAdd(c *gin.Context) *serializer.Response {
	var os role.ProxyService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Add(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// ProxyRemove 移除代理
func ProxyRemove(c *gin.Context) *serializer.Response {
	var os role.ProxyService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Remove(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}
