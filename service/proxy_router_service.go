package service

import (
	"singo/model"
	"singo/serializer"
)

// RouterRegisterService 路由注册服务
type RouterRegisterService struct {
	Path   string `form:"path" json:"path" binding:"required,min=1,max=30"`
	Method string `form:"method" json:"method" binding:"required,min=1,max=30"`
}

// valid 验证表单
func (service *RouterRegisterService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.Router{}).Where("path = ? and method = ?", service.Path, service.Method).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "路由已存在",
		}
	}

	return nil
}

func (service RouterRegisterService) Register() serializer.Response {

	if err := service.valid(); err != nil {
		return *err
	}

	router := model.Router{
		Path:   service.Path,
		Method: service.Method,
	}

	if err := model.DB.Create(&router).Error; err != nil {
		return serializer.Response{Code: 40001, Msg: err.Error()}
	}

	return serializer.Response{
		Data: router,
	}
}
