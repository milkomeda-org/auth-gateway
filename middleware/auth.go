package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"singo/model"
	"singo/serializer"
	"singo/util"
	"strconv"
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
		util.Log().Error("Get user from jwt error: %v", token)
		return nil
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	} else {
		return nil
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getIdFromClaims(c.GetHeader("Authorization"))
		if nil != token && serializer.IsLegal(token) {
			for s := range token {
				c.Set(s, token[s])
			}
			c.Next()
		} else {
			c.JSON(200, serializer.CheckLogin())
			c.Abort()
		}
	}
}

// ResourcePermission 资源权限
func ResourceAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		act := c.Request.Method
		obj := c.Request.URL.RequestURI()
		roles, exists := c.Get("roles")
		if exists {
			rs := roles.([]interface{})
			for _, value := range rs {
				if v, ok := value.(float64); ok {
					sub := strconv.FormatFloat(v, 'f', -1, 64)
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
