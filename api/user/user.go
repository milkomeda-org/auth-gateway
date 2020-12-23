package user

import (
	"auth-gateway/middleware"
	"auth-gateway/serializer"
	user2 "auth-gateway/serializer/user"
	"auth-gateway/service/user"

	"github.com/gin-gonic/gin"
)

// Register 用户注册接口
func Register(c *gin.Context) {
	var us user.RegisterService
	if err := c.ShouldBind(&us); err == nil {
		res := us.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.I18Error(err))
	}
}

// Login 用户登录接口
func Login(c *gin.Context) {
	if u, _ := c.Get("user"); nil != u {
		c.JSON(400, &serializer.Response{
			Code: 0,
			Msg:  "重复登录",
		})
		return
	}
	var us user.LoginService
	if err := c.ShouldBind(&us); err == nil {
		res := us.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.I18Error(err))
	}
}

// Me 用户详情
func Me(c *gin.Context) {
	u, exists := c.Get("user")
	var res *serializer.Response
	if exists {
		res = user2.BuildUserResponse(u.(user2.Session))
	} else {
		user2.BuildUserResponse(*middleware.CurrentUser(c))
	}
	c.JSON(200, res)
}

// Logout 用户登出
func Logout(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

// RoleAdd 添加角色
func RoleAdd(c *gin.Context) *serializer.Response {
	var os user.RoleService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Add(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// RoleRemove 移除角色
func RoleRemove(c *gin.Context) *serializer.Response {
	var os user.RoleService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Remove(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// GroupAdd 添加组
func GroupAdd(c *gin.Context) *serializer.Response {
	var os user.GroupService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Add(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// GroupRemove 移除组
func GroupRemove(c *gin.Context) *serializer.Response {
	var os user.GroupService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Remove(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// OfficeAdd 添加组织
func OfficeAdd(c *gin.Context) *serializer.Response {
	var os user.OfficeService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Add(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// OfficeRemove 移除组织
func OfficeRemove(c *gin.Context) *serializer.Response {
	var os user.OfficeService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Remove(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// PositionAdd 添加职位
func PositionAdd(c *gin.Context) *serializer.Response {
	var os user.PositionService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Add(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}

// PositionRemove 移除职位
func PositionRemove(c *gin.Context) *serializer.Response {
	var os user.PositionService
	if err := c.ShouldBind(&os); err != nil {
		return serializer.I18Error(err)
	}
	if e := os.Remove(); e != nil {
		return serializer.Failed(e)
	}
	return serializer.Success(true)
}
