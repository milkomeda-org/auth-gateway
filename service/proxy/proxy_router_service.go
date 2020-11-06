package proxy

import (
	"oa-auth/initializer/db"
	"oa-auth/model/resource"
	"oa-auth/serializer"
)

// RegisterService 路由注册服务
type RegisterService struct {
	Path   string `form:"path" json:"path" binding:"required,min=1,max=30"`
	Method string `form:"method" json:"method" binding:"required,min=1,max=30"`
}

// valid 验证表单
func (service *RegisterService) valid() *serializer.Response {
	count := 0
	db.DB.Model(&resource.Proxy{}).Where("path = ? and method = ?", service.Path, service.Method).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "路由已存在",
		}
	}

	return nil
}

// Register 注册路由
func (service RegisterService) Register() serializer.Response {

	if err := service.valid(); err != nil {
		return *err
	}

	proxy := resource.Proxy{
		Path:   service.Path,
		Method: service.Method,
	}

	if err := db.DB.Create(&proxy).Error; err != nil {
		return serializer.Response{Code: 40001, Msg: err.Error()}
	}

	return serializer.Response{
		Data: proxy,
	}
}
