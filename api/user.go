package api

import (
	"goa/middleware"
	"goa/serializer"
	"goa/service"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var us service.UserRegisterService
	if err := c.ShouldBind(&us); err == nil {
		res := us.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	if u, _ := c.Get("user"); nil != u {
		c.JSON(400, serializer.Response{
			Code: 0,
			Msg:  "重复登录",
		})
		return
	}
	var us service.UserLoginService
	if err := c.ShouldBind(&us); err == nil {
		res := us.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user, exists := c.Get("user")
	var res serializer.Response
	if exists {
		res = serializer.BuildUserResponse(user.(serializer.UserSession))
	} else {
		serializer.BuildUserResponse(*middleware.CurrentUser(c))
	}
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
