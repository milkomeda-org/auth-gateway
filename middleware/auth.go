package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goa/model"
	"goa/serializer"
	"goa/util"
	"os"
	"strconv"
)

func getIDFromClaims(key string) map[string]interface{} {
	token, err := jwt.Parse(key, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("USER_AUTH_SECRET_KEY")), nil
	})
	if nil != err {
		util.Error("Get user from jwt error: %v", err)
		return nil
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	}
	return nil
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := getIDFromClaims(c.GetHeader("Authorization"))
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

// ResourceAccess 资源授权
func ResourceAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		act := c.Request.Method
		obj := c.Request.URL.RequestURI()
		roles, exists := c.Get("roles")
		if exists && nil != roles {
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
