package user

import (
	"oa-auth/initializer/db"
	user2 "oa-auth/model/user"
	"oa-auth/serializer"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"username" json:"username" binding:"required,min=1,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user user2.User

	if err := db.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return serializer.ParamErr("账号或密码错误", nil)
	}

	if !user.CheckPassword(service.Password) {
		return serializer.ParamErr("账号或密码错误", nil)
	}
	var rs = make(map[int]string, 0)
	// 查询用户继承的角色和自身角色
	rows, _ := db.DB.Raw(`select b.id id, b.alias alias from position_role_mappings a left join roles b on a.role_id = b.id where a.position_id in (select position_id from users where id = ?)`, user.ID).Rows()
	{
		defer rows.Close()
		for rows.Next() {
			var id int
			var alias string
			_ = rows.Scan(&id, &alias)
			rs[id] = alias
		}
	}

	//生成jwt
	token := jwt.New(jwt.SigningMethodHS512)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24*7)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user_name"] = user.UserName
	claims["user_id"] = user.ID
	claims["nick_name"] = user.Nickname
	claims["avatar"] = user.Avatar
	claims["roles"] = &rs
	token.Claims = claims

	tokenStr, err := token.SignedString([]byte(os.Getenv("USER_AUTH_SECRET_KEY")))
	if err != nil {
		return serializer.ParamErr("系统错误", err)
	}

	return serializer.Response{Data: tokenStr}
}

func Exists(userName string) bool {
	count := 0
	db.DB.Model(&user2.User{}).Where("user_name = ?", userName).Count(&count)
	return count > 0
}
