package user

import (
	"oa-auth/middleware"
	"oa-auth/serializer"
	user2 "oa-auth/serializer/user"
	"oa-auth/service/user"

	"github.com/gin-gonic/gin"
)

// Register 用户注册接口
func Register(c *gin.Context) {
	var us user.UserRegisterService
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
		c.JSON(400, serializer.Response{
			Code: 0,
			Msg:  "重复登录",
		})
		return
	}
	var us user.UserLoginService
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
	var res serializer.Response
	if exists {
		res = user2.BuildUserResponse(u.(user2.UserSession))
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
