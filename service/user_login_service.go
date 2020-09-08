package service

import (
	"goa/model"
	"goa/serializer"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	var urm []model.UserRoleMapping
	var rs []int
	model.DB.Model(model.UserRoleMapping{UserID: user.ID}).Find(&urm)
	if nil != urm {
		for i := range urm {
			rs = append(rs, int(urm[i].RoleID))
		}
	}

	//生成jwt
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user_name"] = user.UserName
	claims["user_id"] = user.ID
	claims["nick_name"] = user.Nickname
	claims["avatar"] = user.Avatar
	claims["roles"] = rs
	token.Claims = claims

	tokenStr, err := token.SignedString([]byte(os.Getenv("USER_AUTH_SECRET_KEY")))
	if err != nil {
		return serializer.ParamErr("系统错误", err)
	}

	return serializer.Response{Data: tokenStr}
}
