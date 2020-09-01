package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"singo/model"
	"singo/serializer"
	"singo/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getIdFromClaims(key string) map[string]interface{} {
	token, err := jwt.Parse(key, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("USER_AUTH_SECRET_KEY")), nil
	})
	if nil != err {
		util.Log().Error("Get user from jwt error: %v", err)
		return nil
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	} else {
		return nil
	}
}

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getIdFromClaims(c.GetHeader("Authorization"))
		if nil != token && serializer.IsLegal(token) {
			user, err := model.GetUser(token["user_id"])
			if err == nil {
				c.Set("user", &user)
				c.Set("user_id", token["user_id"])
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}
		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}

// ResourcePermission 资源权限
func ResourceAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		act := c.Request.Method
		obj := c.Request.URL.RequestURI()
		var us []model.UserRoleMapping
		user, _ := c.Get("user")
		if value, ok := user.(*model.User); ok {
			model.DB.Model(&model.UserRoleMapping{UserId: value.ID}).Find(&us)
			if nil != us {
				for i := range us {
					sub := strconv.Itoa(int(us[i].RoleId))
					if ok := model.Enforcer.Enforce(sub, act, obj); ok {
						c.Next()
						return
					}
				}
			}
		}
		c.JSON(200, serializer.NoAccess())
		c.Abort()
	}
}
