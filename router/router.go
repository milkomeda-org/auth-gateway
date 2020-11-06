package router

import (
	"errors"
	"oa-auth/api"
	access2 "oa-auth/api/access"
	group2 "oa-auth/api/group"
	office2 "oa-auth/api/office"
	position2 "oa-auth/api/position"
	aProxy "oa-auth/api/proxy"
	role2 "oa-auth/api/role"
	user2 "oa-auth/api/user"
	"oa-auth/cache"
	"oa-auth/initializer/db"
	"oa-auth/middleware"
	"oa-auth/model/resource"
	"oa-auth/proxy"
	"oa-auth/serializer"
	"oa-auth/util"

	"github.com/gin-gonic/gin"
)

type handlerFunc func(c *gin.Context) serializer.Response

// rest包装器
func restWrapper(execute handlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if p := recover(); nil != p {
				util.Error("panic %s", p)
				c.JSON(500, serializer.Failed(errors.New("严重错误")))
			}
		}()
		c.JSON(200, execute(c))
	}
}

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(util.GinLogger(util.Logger), util.GinRecovery(util.Logger, true))

	// 中间件, 顺序不能改
	r.Use(middleware.Cors())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/login", user2.Login)

		//路由添加
		v1.POST("proxy/register", aProxy.RegisterProxy)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// Auth Routes

			user := auth.Group("user")
			user.GET("me", user2.Me)
			user.DELETE("logout", user2.Logout)

			//需要角色权限的
			access := auth.Group("")
			access.Use(middleware.ResourceAccess())
			{
				// 组织管理
				office := access.Group("office")
				{
					office.POST("", restWrapper(office2.Create))
					office.PUT("", restWrapper(office2.Update))
					office.DELETE("", restWrapper(office2.Delete))
					office.GET("", restWrapper(office2.View))
				}
				// 职位管理
				position := access.Group("/position")
				{
					position.POST("", restWrapper(position2.Create))
					position.PUT("", restWrapper(position2.Update))
					position.DELETE("", restWrapper(position2.Delete))
					position.GET("", restWrapper(position2.View))
					position.POST("/role", restWrapper(position2.RoleAdd))
					position.DELETE("/role", restWrapper(position2.RoleRemove))
				}
				// 用户组管理
				group := access.Group("group")
				{
					group.POST("", restWrapper(group2.Create))
					group.PUT("", restWrapper(group2.Update))
					group.DELETE("", restWrapper(group2.Delete))
					group.GET("", restWrapper(group2.View))
				}

				// 用户注册
				access.POST("user/register", user2.Register)

				//角色管理
				role := access.Group("role")
				{
					role.POST("", role2.Create)
					role.DELETE("", role2.Delete)
				}

				//授权管理
				//x-www sub act obj
				access.POST("permission", restWrapper(access2.Add))
				// query sub act obj
				access.DELETE("permission", restWrapper(access2.Remove))
			}
		}

		// 路由代理
		appProxy := v1.Group("proxy")
		appProxy.Use(middleware.AuthRequired())
		appProxy.Use(middleware.ResourceAccess())
		{
			// 人员信息表
			var rs []resource.Proxy
			db.DB.Model(resource.Proxy{}).Find(&rs)
			if nil != rs {
				for _, v := range rs {
					appProxy.Handle(v.Method, v.Path, proxy.HostProxy)
				}
			}
			cache.AppProxy = appProxy
		}
	}
	return r
}
