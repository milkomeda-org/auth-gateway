package middleware

import (
	"auth-gateway/initializer/db"
	"auth-gateway/serializer"
	jwt2 "auth-gateway/serializer/jwt"
	u2 "auth-gateway/serializer/user"
	"auth-gateway/util/log"
	"fmt"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *u2.Session {
	if user, err := getIDFromClaims(c.GetHeader("Authorization")); nil == err {
		return &user
	} else {
		c.JSON(200, serializer.CheckLogin())
		c.Abort()
		return nil
	}
}

func getIDFromClaims(key string) (u2.Session, error) {
	token, err := jwt.Parse(key, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("USER_AUTH_SECRET_KEY")), nil
	})
	if nil != err {
		log.Error("Get user from jwt error: %v", err)
		return u2.Session{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if nil != claims && jwt2.IsLegal(claims) {
			var us = u2.Session{}
			us.UserName = claims["user_name"].(string)
			us.Avatar = claims["avatar"].(string)
			us.NickName = claims["nick_name"].(string)
			uidF := claims["user_id"]
			if v, ok := uidF.(float64); ok {
				uidS := strconv.FormatFloat(v, 'f', -1, 64)
				if uid, err := strconv.ParseInt(uidS, 0, 64); nil == err {
					us.UserID = int(uid)
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
	return u2.Session{}, nil
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
			userSession := user.(u2.Session)
			rs := userSession.Roles
			if userSession.UserName == "admin" {
				c.Next()
				return
			}
			for rid, _ := range rs {
				sub := strconv.Itoa(rid)
				if ok := db.Enforcer.Enforce(sub, act, obj); ok {
					c.Next()
					return
				}
			}
		}
		c.JSON(200, serializer.NoAccess())
		c.Abort()
	}
}
