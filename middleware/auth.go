package middleware

import (
	"fmt"
	"oa-auth/initializer"
	"oa-auth/serializer"
	"oa-auth/util"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *serializer.UserSession {
	if user, err := getIDFromClaims(c.GetHeader("Authorization")); nil == err {
		return &user
	} else {
		c.JSON(200, serializer.CheckLogin())
		c.Abort()
		return nil
	}
}

func getIDFromClaims(key string) (serializer.UserSession, error) {
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
		return serializer.UserSession{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if nil != claims && serializer.IsLegal(claims) {
			var us = serializer.UserSession{}
			us.UserName = claims["user_name"].(string)
			us.Avatar = claims["avatar"].(string)
			us.NickName = claims["nick_name"].(string)
			uidF := claims["user_id"]
			if v, ok := uidF.(float64); ok {
				uidS := strconv.FormatFloat(v, 'f', -1, 64)
				if uid, err := strconv.ParseUint(uidS, 0, 64); nil == err {
					us.UserID = uint(uid)
				}
			}
			rs := claims["roles"].(map[string]interface{})
			rsSource := make(map[int]string)
			for k, v := range rs {
				if i, err := strconv.Atoi(k); nil == err {
					if j, ok := v.(string); ok {
						rsSource[i] = j
					}
				}
			}
			us.Roles = rsSource
			return us, nil
		}
	}
	return serializer.UserSession{}, nil
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if session, err := getIDFromClaims(c.GetHeader("Authorization")); nil == err {
			c.Set("user", session)
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
		obj := c.Request.URL.Path
		user, exists := c.Get("user")
		if exists && nil != user {
			rs := user.(serializer.UserSession).Roles
			for k, _ := range rs {
				sub := strconv.Itoa(k)
				if ok := initializer.Enforcer.Enforce(sub, act, obj); ok {
					c.Next()
					return
				}
			}
		}
		c.JSON(200, serializer.NoAccess())
		c.Abort()
	}
}
