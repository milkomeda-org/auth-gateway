package server

import (
	"errors"
	"oa-auth/api"
	"oa-auth/api/organization"
	"oa-auth/api/permission"
	"oa-auth/cache"
	"oa-auth/initializer"
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
		v1.POST("user/login", api.UserLogin)

		//路由添加
		v1.POST("proxy/router/register", api.RegisterRouter)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// Auth Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)

			//需要角色权限的
			access := auth.Group("")
			access.Use(middleware.ResourceAccess())
			{
				// 组织管理
				access.POST("office", restWrapper(organization.OfficeCreate))
				access.PUT("office", restWrapper(organization.OfficeUpdate))
				access.DELETE("office", restWrapper(organization.OfficeDelete))
				access.GET("office", restWrapper(organization.OfficeView))
				// 职位管理
				position := access.Group("/position")
				{
					position.POST("", restWrapper(organization.PositionCreate))
					position.PUT("", restWrapper(organization.PositionUpdate))
					position.DELETE("", restWrapper(organization.PositionDelete))
					position.GET("", restWrapper(organization.PositionView))
					position.POST("/role", restWrapper(organization.PositionRoleAdd))
					position.DELETE("/role", restWrapper(organization.PositionRoleRemove))
				}
				//// 用户组管理
				access.POST("group", restWrapper(organization.GroupCreate))
				access.PUT("group", restWrapper(organization.GroupUpdate))
				access.DELETE("group", restWrapper(organization.GroupDelete))
				access.GET("group", restWrapper(organization.GroupView))

				// 用户注册
				access.GET("user/register", api.UserRegister)

				//角色管理
				access.POST("role", api.CreateRole)
				access.DELETE("role", api.DeleteRole)

				//授权管理
				//x-www sub act obj
				access.POST("permission", restWrapper(permission.AddAccess))
				//query sub act obj
				access.DELETE("permission", restWrapper(permission.RemoveAccess))
			}
		}

		// 路由代理
		appProxy := v1.Group("proxy")
		appProxy.Use(middleware.AuthRequired())
		appProxy.Use(middleware.ResourceAccess())
		{
			// 人员信息表
			var rs []resource.Router
			initializer.DB.Model(resource.Router{}).Find(&rs)
			if nil != rs {
				for _, v := range rs {
					appProxy.Handle(v.Method, v.Path, proxy.HostProxy)
				}
			}
			cache.AppProxyRouter = appProxy
		}
	}
	return r
}
